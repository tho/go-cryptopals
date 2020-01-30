package challenge3

import (
	"math"
	"strings"
)

// FindSingleByteXORKey finds the byte which was used to encrypt the buffer s.
//
// The key-byte is determined by decrypting the buffer using all byte values in
// sequence and scoring the results by character frequency. The key and its
// score are returned.
func FindSingleByteXORKey(s []byte) (key byte, score int) {
	var maxScore int

	for i := 0; i <= math.MaxUint8; i++ {
		xor := ComputeSingleByteXOR(s, byte(i))
		score := scoreByCharacterFrequency(xor)

		if score > maxScore {
			key = byte(i)
			maxScore = score
		}
	}

	return key, maxScore
}

// ComputeSingleByteXOR computes the XOR combination of slice s and byte c.
func ComputeSingleByteXOR(s []byte, c byte) []byte {
	xor := make([]byte, len(s))
	for i := range xor {
		xor[i] = s[i] ^ c
	}

	return xor
}

func scoreByCharacterFrequency(s []byte) int {
	const relativeCharFreqReversed = "UuLlDdRrHhSs NnIiOoAaTtEe"

	var score int

	for _, c := range s {
		// IndexByte returns -1 if c is not present in the string. To
		// account for this, 1 is added to idx to effectively ignore
		// characters which are not in relativeCharFreqReserved.
		idx := strings.IndexByte(relativeCharFreqReversed, c)
		score += idx + 1
	}

	return score
}
