package challenge3

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func Example() {
	const hexStr = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	ciphertext, _ := hex.DecodeString(hexStr)
	key, _ := FindSingleByteXORKey(ciphertext)
	plaintext := ComputeSingleByteXOR(ciphertext, key)

	fmt.Printf("%s\n", plaintext)
	// Output: Cooking MC's like a pound of bacon
}

func TestComputeSingleByteXOR(t *testing.T) {
	tests := map[string]struct {
		s    []byte
		c    byte
		want []byte
	}{
		"valid": {
			s:    []byte{0x00, 0x00},
			c:    0xff,
			want: []byte{0xff, 0xff},
		},
		"nil s": {
			s:    nil,
			c:    0xff,
			want: []byte{},
		},
		"blank s": {
			s:    []byte{},
			c:    0xff,
			want: []byte{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := append(tt.s[:0:0], tt.s...)

			got := ComputeSingleByteXOR(s, tt.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeSingleByteXOR(%v, %v) = %v, want: %v", tt.s, tt.c, got, tt.want)
			}
			if !reflect.DeepEqual(s, tt.s) {
				t.Errorf("computeSingleByteXOR(%v, %v) changed input buffer to %v", tt.s, tt.c, s)
			}
		})
	}
}
