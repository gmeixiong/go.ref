// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bcrypter defines the mechanisms for blessings based
// encryption and decryption.
package bcrypter

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"sync"

	"v.io/v23/context"
	"v.io/v23/security"

	"v.io/x/lib/ibe"
)

const hashTruncation = 16

// Crypter provides operations for encrypting and decrypting messages for
// principals with specific blessings.
//
// In particular, it offers a mechanism to encrypt a message for a specific
// blessing pattern so that it can only be decrypted by crypters that possess
// a private key for a blessing matched by that pattern. Such a private key
// is generated by the identity provider that granted the blessing.
type Crypter struct {
	mu sync.RWMutex
	// root blessing -> []ibe.Params
	params map[string][]ibe.Params
	// paramsId -> patternId -> ibe.PrivateKey
	keys map[string]map[string]ibe.PrivateKey
}

// Ciphertext represents the ciphertext generated by a Crypter.
type Ciphertext struct {
	wire WireCiphertext
}

// Root represents an identity provider for the purposes of blessings based
// encryption and decryption.
//
// It generates private keys for specific blessings which can be used
// to decrypt any message encrypted for a pattern matched by the blessing (
// assuming the encryption used this identity provider's parameters).
type Root struct {
	// master is the IBE Master that this root uses to extract IBE
	// private keys.
	master ibe.Master
	// blessing is the blessing name of the identity provider. The identity
	// provider  can extract private keys for blessings that are extensions
	// of this blessing name.
	blessing string
}

// Params represents the public parameters of an identity provider (aka Root).
type Params struct {
	// blessing is the blessing name of the identity provider.
	blessing string
	// params are the public IBE params of the identity provider
	params ibe.Params
}

// PrivateKey represent the private key corresponding to a blessing.
//
// The private key can be used for decrypting any message encrypted using
// a pattern matched by the blessing (assuming the private key and encryption
// used the same identity provider parameters).
type PrivateKey struct {
	// blessing is the blessing for which this private key was extracted for.
	blessing string
	// params represents the public parameters of the identity provider
	// that extracted this private key. The blessing must be an extension
	// of params.blessing.
	params Params
	// keys contain private keys extracted for each blessing pattern that is
	// matched by the blessing and is an extension of root.blessing.
	//
	// For example, if the blessing is "google/u/alice/phone" and root.blessing
	// is "google/u" then the keys are extracted for patterns "google/u",
	// "google/u/alice", "google/u/alice/phone", and "google/u/alice/phone/$".
	//
	// The private keys are listed in increasing order of the lengths of the
	// corresponding patterns.
	keys []ibe.PrivateKey
}

// Encrypt encrypts the provided fixed-length 'plaintext' so that it can
// only be decrypted by a crypter possessing a private key for a blessing
// matching the provided blessing pattern.
//
// Encryption makes use of the public parameters of the identity provider
// that is authoritative on the set of blessings that match the provided
// blessing pattern. These paramaters must have been previously added to
// this crypter via AddParams.
func (c *Crypter) Encrypt(ctx *context.T, forPattern security.BlessingPattern, plaintext *[32]byte) (*Ciphertext, error) {
	if !forPattern.IsValid() {
		return nil, fmt.Errorf("provided blessing pattern %v is invalid", forPattern)
	}
	ciphertext := &Ciphertext{wire: WireCiphertext{PatternId: idPattern(forPattern), Bytes: make(map[string][]byte)}}
	paramsFound := false
	c.mu.RLock()
	defer c.mu.RUnlock()
	for name, ibeParamsList := range c.params {
		if !isExtensionOf(forPattern, name) {
			continue
		}
		for _, ibeParams := range ibeParamsList {
			ctxt := make([]byte, ibe.CiphertextSize)
			if err := ibeParams.Encrypt(string(forPattern), (*plaintext)[:], ctxt); err != nil {
				return nil, NewErrInternal(ctx, err)
			}
			paramsId, err := idParams(ibeParams)
			if err != nil {
				return nil, NewErrInternal(ctx, err)
			}
			paramsFound = true
			ciphertext.wire.Bytes[paramsId] = ctxt
		}
	}
	if !paramsFound {
		return nil, NewErrNoParams(ctx, forPattern)
	}
	return ciphertext, nil
}

// Decrypt decrypts the provided 'ciphertext' and returns the corresponding
// plaintext.
//
// Decryption succeeds only if this crypter possesses a private key for a
// blessing that matches the blessing pattern corresponding to the ciphertext.
func (c *Crypter) Decrypt(ctx *context.T, ciphertext *Ciphertext) (*[32]byte, error) {
	var (
		plaintext [32]byte
		keyFound  bool
	)
	c.mu.RLock()
	defer c.mu.RUnlock()
	for paramsId, cbytes := range ciphertext.wire.Bytes {
		if keys, found := c.keys[paramsId]; !found {
			continue
		} else if key, found := keys[ciphertext.wire.PatternId]; !found {
			continue
		} else if err := key.Decrypt(cbytes, plaintext[:]); err != nil {
			return nil, NewErrInternal(ctx, err)
		}
		keyFound = true
		break
	}
	if !keyFound {
		return nil, NewErrPrivateKeyNotFound(ctx)
	}
	return &plaintext, nil
}

// AddKey adds the provided private key 'key' and the associated public
// parameters (key.Params()) to this crypter.
func (c *Crypter) AddKey(ctx *context.T, key *PrivateKey) error {
	patterns := matchedBy(key.blessing, key.params.blessing)
	if got, want := len(key.keys), len(patterns); got != want {
		return NewErrInvalidPrivateKey(ctx, fmt.Errorf("got %d IBE private keys for blessing %v (and root blessing %v), expected %d", got, key.blessing, key.params.blessing, want))
	}

	paramsId, err := idParams(key.params.params)
	if err != nil {
		return NewErrInternal(ctx, err)
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.params[key.params.blessing] = append(c.params[key.params.blessing], key.params.params)
	if _, found := c.keys[paramsId]; !found {
		c.keys[paramsId] = make(map[string]ibe.PrivateKey)
	}
	for i, p := range patterns {
		c.keys[paramsId][idPattern(p)] = key.keys[i]
	}
	return nil
}

// AddParams adds the provided identity provider parameters to this crypter.
//
// The added parameters would be used to encrypt plaintexts for blessing patterns
// that the identity provider is authoritative on.
func (c *Crypter) AddParams(ctx *context.T, params Params) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	// TODO(ataly, ashankar): Avoid adding duplicate params to the list.
	c.params[params.blessing] = append(c.params[params.blessing], params.params)
	return nil
}

// Blessing returns the blessing that this private key was extracted for.
func (k *PrivateKey) Blessing() string {
	return k.blessing
}

// Params returns the public parameters of the identity provider that
// extracted this private key.
func (k *PrivateKey) Params() Params {
	return k.params
}

// Params returns the public parameters of the identity provider represented
// by 'r'.
func (r *Root) Params() Params {
	return Params{blessing: r.blessing, params: r.master.Params()}
}

// Extract returns a private key for the provided blessing.
//
// The private key can be used for decrypting any message encrypted using a
// pattern matched by the blessing (assuming the encryption made use of the
// public parameters of this root).
func (r *Root) Extract(ctx *context.T, blessing string) (*PrivateKey, error) {
	patterns := matchedBy(blessing, r.blessing)
	if len(patterns) == 0 {
		return nil, fmt.Errorf("blessing %v does not match the blessing pattern this root is authoritative on: %v", blessing, r.blessing)
	}
	key := &PrivateKey{
		blessing: blessing,
		params:   r.Params(),
		keys:     make([]ibe.PrivateKey, len(patterns)),
	}
	for i, p := range patterns {
		ibeKey, err := r.master.Extract(string(p))
		if err != nil {
			return nil, NewErrInternal(ctx, err)
		}
		key.keys[i] = ibeKey
	}
	return key, nil
}

// Blessing returns the blessing name of the identity provider with
// public parameters 'p'.
func (p *Params) Blessing() string {
	return p.blessing
}

// NewCrypter returns a new Crypter with an empty set of private keys
// and identity provider parameters.
func NewCrypter() *Crypter {
	return &Crypter{params: make(map[string][]ibe.Params), keys: make(map[string]map[string]ibe.PrivateKey)}
}

// NewRoot returns a new root identity provider that has the provided
// blessing name and uses the provided 'master' for setting up identity-based
// encryption.
func NewRoot(blessing string, master ibe.Master) *Root {
	return &Root{blessing: blessing, master: master}
}

// matchedBy returns the set of blessing patterns (in increasing order
// of length) that are matched by the provided 'blessing' and are equal
// to or extensions of the blessing name 'root'.
func matchedBy(blessing, root string) []security.BlessingPattern {
	if !security.BlessingPattern(root).MatchedBy(blessing) {
		return nil
	}
	patterns := make([]security.BlessingPattern, strings.Count(blessing, security.ChainSeparator)+2-strings.Count(string(root), security.ChainSeparator))
	patterns[len(patterns)-1] = security.BlessingPattern(blessing).MakeNonExtendable()
	patterns[len(patterns)-2] = security.BlessingPattern(blessing)
	for idx := len(patterns) - 3; idx >= 0; idx-- {
		blessing = blessing[0:strings.LastIndex(blessing, string(security.ChainSeparator))]
		patterns[idx] = security.BlessingPattern(blessing)
	}
	return patterns
}

// idPattern returns a 128-bit truncated SHA-256 hash of a blessing pattern.
func idPattern(pattern security.BlessingPattern) string {
	h := sha256.Sum256([]byte(pattern))
	truncated := h[:hashTruncation]
	return string(truncated)
}

// idParams returns a 128-bit truncated SHA-256 hash of the marshaled IBE params.
func idParams(params ibe.Params) (string, error) {
	paramsBytes, err := ibe.MarshalParams(params)
	if err != nil {
		return "", err
	}
	h := sha256.Sum256(paramsBytes)
	truncated := h[:hashTruncation]
	return string(truncated), nil
}

// isExtensionOf returns true if the all blessings matching the provided
// 'pattern' are an extension of the provided 'root' blessing
func isExtensionOf(pattern security.BlessingPattern, root string) bool {
	return string(pattern) == root || strings.HasPrefix(string(pattern), root+security.ChainSeparator)
}
