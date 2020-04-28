package ecdh

import (
	"crypto"
)

// ECDH - The main interface ECDH.
type ECDH interface {
	GenerateKeyPair() (crypto.PrivateKey, crypto.PublicKey, error)
	Marshal(crypto.PublicKey) []byte
	MarshalSK(crypto.PrivateKey) []byte
	Unmarshal([]byte) (crypto.PublicKey, bool)
	UnmarshalSK([]byte) (crypto.PrivateKey, bool)
	GenerateSharedSecret(crypto.PrivateKey, crypto.PublicKey) ([]byte, error)
}
