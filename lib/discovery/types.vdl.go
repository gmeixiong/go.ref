// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package discovery

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/discovery"
)

type EncryptionAlgorithm int32

func (EncryptionAlgorithm) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/discovery.EncryptionAlgorithm"`
}) {
}

type EncryptionKey []byte

func (EncryptionKey) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/discovery.EncryptionKey"`
}) {
}

type Uuid []byte

func (Uuid) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/discovery.Uuid"`
}) {
}

// Advertisement holds a set of service properties to advertise.
type Advertisement struct {
	// The service to advertise.
	Service discovery.Service
	// Type of encryption applied to the advertisement so that it can
	// only be decoded by authorized principals.
	EncryptionAlgorithm EncryptionAlgorithm
	// If the advertisement is encrypted, then the data required to
	// decrypt it. The format of this data is a function of the algorithm.
	EncryptionKeys []EncryptionKey
	// Hash of the current advertisement.
	Hash []byte
	// The addresses (vanadium object names) that the advertisement directory service
	// is served on. See directory.vdl.
	DirAddrs []string
	// TODO(jhahn): Add proximity.
	// TODO(jhahn): Use proximity for Lost.
	Lost bool
}

func (Advertisement) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/discovery.Advertisement"`
}) {
}

func init() {
	vdl.Register((*EncryptionAlgorithm)(nil))
	vdl.Register((*EncryptionKey)(nil))
	vdl.Register((*Uuid)(nil))
	vdl.Register((*Advertisement)(nil))
}

const NoEncryption = EncryptionAlgorithm(0)

const TestEncryption = EncryptionAlgorithm(1)

const IbeEncryption = EncryptionAlgorithm(2)
