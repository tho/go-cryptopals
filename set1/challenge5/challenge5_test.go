package challenge5

import (
	"fmt"
	"reflect"
	"testing"
)

func Example() {
	var (
		plaintext = []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
		key = []byte("ICE")
	)

	ciphertext := ComputeRepeatingKeyXOR(plaintext, key)

	fmt.Printf("%x\n", ciphertext)
	// Output: 0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f
}

func TestComputeRepeatingKeyXOR(t *testing.T) {
	tests := map[string]struct {
		s    []byte
		key  []byte
		want []byte
	}{
		"valid": {
			s:    []byte{0x00, 0x00, 0x00},
			key:  []byte{0xff},
			want: []byte{0xff, 0xff, 0xff},
		},
		"nil s": {
			s:    nil,
			key:  []byte{0xff},
			want: []byte{},
		},
		"blank s": {
			s:    []byte{},
			key:  []byte{0xff},
			want: []byte{},
		},
		"nil key": {
			s:    []byte{0x00, 0x00, 0x00},
			key:  nil,
			want: []byte{0x00, 0x00, 0x00},
		},
		"blank key": {
			s:    []byte{0x00, 0x00, 0x00},
			key:  []byte{},
			want: []byte{0x00, 0x00, 0x00},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := append(tt.s[:0:0], tt.s...)

			got := ComputeRepeatingKeyXOR(s, tt.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeRepeatingKeyXOR(%v, %v) = %v, want: %v", tt.s, tt.key, got, tt.want)
			}
			if !reflect.DeepEqual(s, tt.s) {
				t.Errorf("computeRepeatingKeyXOR(%v, %v) changed input buffer to %v", tt.s, tt.key, s)
			}
		})
	}
}
