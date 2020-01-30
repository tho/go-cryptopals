package challenge8

import (
	"github.com/tho/go-cryptopals/set1/challenge6"
)

// IsECB returns true if s has been encrypted with ECB.
func IsECB(s []byte) bool {
	seen := make(map[string]bool)

	for _, x := range challenge6.SplitSize(s, 16) {
		if seen[string(x)] {
			return true
		}

		seen[string(x)] = true
	}

	return false
}
