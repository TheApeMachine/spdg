package spdg

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
)

/*
Layer is the lowest level of the core SPDG model. It takes in byte buffers
which are commonly known as the "last known datatype" in the SPDG model.
That means it is the only other datatype next to SPDGs themselves that
can travel parts of the pipeline. Any ingress point's main concern should
be to get to SPDG as quick as possible by having a contained scope with a
layer of any to byte buffer converters catch any incoming types, and from
convert to go to SPDGs. Only then should real scope be granted.
The following features are important:

- Encrypted   : Encryption is a first class citizen in SPDG and a clear
  =========
	policy is that data should not leak. That means multiple encodes and
	decodes as data flows through the pipeline. No public method should
	ever return or accept unencrypted packets.

- Private     : The type follows the strictest privacy guidelines where
  =======
	only access to the correct key combination values can be inspected
	or written. Access is granular to the packet level, so a validated
	connection achieves nothing, and is therefor not needed. This keeps
	the data type future proofed and inline with ever more complex
	privacy regulations such as GDPR and the likes.

- Destructive : For reasons of tamper-proofing the type has the option
  ===========
	to layer its own public key, after which the data is inaccessible.
	Another benefit is that is gives the data an inherent value besides
	its external value as a secure and trusted token of information.
	It is important to note that the data is not lost in this case.
	Canonically it exists and knowledge of it reverberates through the
	objects it has interacted with. But it was considered void of value
	and thus marked as non-grata. This will combat value inflation.
*/
type Layer interface {
	State() (Status, Reason)
	Peek(bytes.Buffer)
	Poke(bytes.Buffer)
}

/* Layer defines the interface for the common way data is stored. Encryption and anonimity
are first class citizens in Layers. */
type ProtoLayer struct {
	key *rsa.PublicKey
	sec *rsa.PrivateKey
	dat []byte
	err error
}

func NewProtoLayer(ctx context.Context) Layer {
	return ProtoLayer{}
}

func (layer ProtoLayer) State() (Status, Reason) {
	return OK, BUSY
}

func (layer ProtoLayer) Peek(dat bytes.Buffer) {
	layer.sec, layer.err = rsa.GenerateKey(rand.Reader, 2048)
	layer.key = &layer.sec.PublicKey
	layer.seal(dat)
}

func (layer ProtoLayer) Poke(dat bytes.Buffer) {
}

func (layer ProtoLayer) seal(dat bytes.Buffer) {
	layer.dat, layer.err = rsa.EncryptOAEP(
		sha512.New(),
		rand.Reader,
		layer.key,
		dat.Bytes(),
		nil,
	)
}
