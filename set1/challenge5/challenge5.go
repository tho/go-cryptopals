package challenge5

// ComputeRepeatingKeyXOR sequentially applies each byte of the key. The first
// byte of key will be XOR'd against the first byte of s, the second byte of key
// against the second byte of s, and so on. After the last byte of key has been
// XOR'd, the first byte is used again.
//
// ComputeRepeatingKeyXOR returns s if the length of key is 0.
func ComputeRepeatingKeyXOR(s, key []byte) []byte {
	if len(key) == 0 {
		return s
	}

	xor := make([]byte, len(s))
	for i := range xor {
		xor[i] = s[i] ^ key[i%len(key)]
	}

	return xor
}
