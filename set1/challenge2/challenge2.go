package challenge2

import (
	"errors"
)

// ComputeFixedXOR computes the XOR combination of two equal length buffers.
//
// ComputeFixedXOR returns an error if the lengths of s1 and s2 are unequal.
func ComputeFixedXOR(s1, s2 []byte) ([]byte, error) {
	if len(s1) != len(s2) {
		return nil, errors.New("unequal length buffers")
	}

	xor := make([]byte, len(s1))
	for i := range xor {
		xor[i] = s1[i] ^ s2[i]
	}

	return xor, nil
}
