package bls

/*
#cgo CFLAGS:-I ${SRCDIR}../contrib/relic/include
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS:-lstdc++
#include "bls_cgo.hpp"
*/
import "C"
import "unsafe"

// PrivateKey --
type PrivateKey struct{}

func (priv *PrivateKey) getPointer() (p *C.blsPrivateKey) {
	// #nosec
	return (*C.blsPrivateKey)(unsafe.Pointer(priv))
}

// SetFromBytes --
func (priv *PrivateKey) SetFromBytes(data []byte) {
	C.blsPrivateKeyFromBytes(priv.getPointer(), unsafe.Pointer(&data[0]), 1)
}