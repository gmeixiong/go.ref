// Package crypto implements encryption and decryption interfaces intended for
// securing communication over VCs.
package crypto

import "veyron/runtimes/google/lib/iobuf"

type Encrypter interface {
	// Encrypt encrypts the provided plaintext data and returns the
	// corresponding ciphertext slice (or nil if an error is returned).
	//
	// It always calls Release on plaintext and thus plaintext should not
	// be used after calling Encrypt.
	Encrypt(plaintext *iobuf.Slice) (ciphertext *iobuf.Slice, err error)
}

type Decrypter interface {
	// Decrypt decrypts the provided ciphertext slice and returns the
	// corresponding plaintext (or nil if an error is returned).
	//
	// It always calls Release on ciphertext and thus ciphertext should not
	// be used after calling Decrypt.
	Decrypt(ciphertext *iobuf.Slice) (plaintext *iobuf.Slice, err error)
}

type Crypter interface {
	Encrypter
	Decrypter
	String() string
}
