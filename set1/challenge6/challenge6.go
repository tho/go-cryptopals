package challenge6

import (
	"errors"
	"math/bits"

	"github.com/tho/go-cryptopals/set1/challenge3"
)

// FindRepeatingKeyXORKey finds the key which was used to encrypt the buffer s.
// The key length is assumed to be in range [minKeySize, maxKeySize].
func FindRepeatingKeyXORKey(s []byte, minKeySize, maxKeySize int) ([]byte, error) {
	keySize, err := findRepeatingKeyXORKeySize(s, minKeySize, maxKeySize)
	if err != nil {
		return nil, err
	}

	blocks := SplitSize(s, keySize)
	transposedBlocks := transpose(blocks)

	key := make([]byte, len(transposedBlocks))
	for i, block := range transposedBlocks {
		key[i], _ = challenge3.FindSingleByteXORKey(block)
	}

	return key, nil
}

func findRepeatingKeyXORKeySize(s []byte, minKeySize, maxKeySize int) (int, error) {
	var keySize int

	minScore := -1.0

	for size := minKeySize; size <= maxKeySize; size++ {
		score, err := scoreByHammingDistance(s, size)
		if err != nil {
			return 0, err
		}

		if score < minScore || minScore == -1 {
			keySize = size
			minScore = score
		}
	}

	return keySize, nil
}

func scoreByHammingDistance(s []byte, keySize int) (float64, error) {
	// blocks are processed in pairs. Adjust nblocks such that the last
	// block is ignored if nblocks would be uneven.
	nblocks := len(s) / keySize
	nblocks -= nblocks % 2

	if nblocks < 2 {
		return 0, errors.New("invalid input or key size")
	}

	// If s is not a multiple of keySize*n, the last subslice returned by
	// SplitSizeN contains the unsplit remainder. Add 1 to nblocks such that
	// the nblocks'th block contains a full block as opposed to a full block
	// plus remainder.
	blocks := SplitSizeN(s, keySize, nblocks+1)

	var (
		sum float64
		n   int
	)

	for i := 0; i < nblocks-1; i += 2 {
		d, err := computeHammingDistance(blocks[i], blocks[i+1])
		if err != nil {
			return 0, err
		}

		sum += float64(d) / float64(keySize)
		n++
	}

	return sum / float64(n), nil
}

func computeHammingDistance(s1, s2 []byte) (int, error) {
	if len(s1) != len(s2) {
		return 0, errors.New("unequal length buffers")
	}

	var distance int

	for i := 0; i < len(s1); i++ {
		distance += bits.OnesCount8(s1[i] ^ s2[i])
	}

	return distance, nil
}

func transpose(m [][]byte) [][]byte {
	if m == nil {
		return nil
	}

	var max int

	for _, row := range m {
		if len(row) > max {
			max = len(row)
		}
	}

	transposed := make([][]byte, max)

	for _, row := range m {
		for j, c := range row {
			transposed[j] = append(transposed[j], c)
		}
	}

	return transposed
}

// Generic splitSize: splits s into subslices of length size.
func genSplitSize(s []byte, size, n int) [][]byte {
	if size <= 0 || n == 0 {
		return nil
	}

	// Determine the number of full subslices of length size.
	if n < 0 || n > len(s)/size {
		n = len(s) / size
	}

	// Account for an empty input slice s and the remainder.
	if len(s) == 0 || len(s)%size != 0 {
		n++
	}

	a := make([][]byte, n)

	i := 0
	for i < n-1 {
		a[i] = s[:size]
		s = s[size:]
		i++
	}

	// Save the remainder in the last subslice.
	a[i] = s

	return a
}

// SplitSize slices s into all subslices of length size and returns a slice of
// the subslices.
// If the length of s is not a multiple of size the length of the last subslice
// will be less than size.
// If size is negative or 0 the result is nil (zero subslice)
func SplitSize(s []byte, size int) [][]byte {
	return genSplitSize(s, size, -1)
}

// splitSizeN slices s into subslices of length size and returns a slice of the
// subslices.
// If the length of s is not a multiple of size*n the length of the last subslice
// will be less than size.
// If size is negative or 0 the result is nil (zero subslice).
//
// The count determines the number of subslices to return:
//
//     n > 0: at most n subslices; the last subslice will be the unsplit remainder.
//     n == 0: the result is nil (zero subslices)
//     n < 0: all subslices
func SplitSizeN(s []byte, size, n int) [][]byte {
	return genSplitSize(s, size, n)
}
