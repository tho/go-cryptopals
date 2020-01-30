package challenge4

import (
	"bufio"
	"encoding/hex"
	"io"

	"github.com/tho/go-cryptopals/set1/challenge3"
)

// DetectSingleCharacterXOR reads hexadecimal ciphertexts from r, one per line,
// and detects the one which has been encrypted by single-character XOR.
func DetectSingleCharacterXOR(r io.Reader) ([]byte, error) {
	var (
		plaintext []byte
		maxScore  int
	)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ciphertext, err := hex.DecodeString(scanner.Text())
		if err != nil {
			return nil, err
		}

		key, score := challenge3.FindSingleByteXORKey(ciphertext)
		if score > maxScore {
			plaintext = challenge3.ComputeSingleByteXOR(ciphertext, key)
			maxScore = score
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return plaintext, nil
}
