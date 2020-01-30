package challenge7

import (
	"crypto/cipher"
	"unsafe"
)

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

// ecbDecrypter implements the cipher.BlockMode interface.
type ecbDecrypter ecb

// ecbDecAble is an interface implemented by ciphers that have a specific
// optimized implementation of ECB decryption.
// newECBDecrypter will check for this interface and return the specific
// BlockMode if found.
type ecbDecAble interface {
	NewECBDecrypter() cipher.BlockMode
}

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	if ecb, ok := b.(ecbDecAble); ok {
		return ecb.NewECBDecrypter()
	}

	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int {
	return x.blockSize
}

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("input not full blocks")
	}

	if len(dst) < len(src) {
		panic("output smaller than input")
	}

	if inexactOverlap(dst[:len(src)], src) {
		panic("invalid buffer overlap")
	}

	for len(src) > 0 {
		x.b.Decrypt(dst[:x.blockSize], src[:x.blockSize])

		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

//
// Copy of crypto/internal/subtle follows.
//

// anyOverlap reports whether x and y share memory at any (not necessarily
// corresponding) index. The memory beyond the slice length is ignored.
func anyOverlap(x, y []byte) bool {
	return len(x) > 0 && len(y) > 0 &&
		uintptr(unsafe.Pointer(&x[0])) <= uintptr(unsafe.Pointer(&y[len(y)-1])) &&
		uintptr(unsafe.Pointer(&y[0])) <= uintptr(unsafe.Pointer(&x[len(x)-1]))
}

// inexactOverlap reports whether x and y share memory at any non-corresponding
// index. The memory beyond the slice length is ignored. Note that x and y can
// have different lengths and still not have any inexact overlap.
//
// inexactOverlap can be used to implement the requirements of the crypto/cipher
// AEAD, Block, BlockMode and Stream interfaces.
func inexactOverlap(x, y []byte) bool {
	if len(x) == 0 || len(y) == 0 || &x[0] == &y[0] {
		return false
	}

	return anyOverlap(x, y)
}
