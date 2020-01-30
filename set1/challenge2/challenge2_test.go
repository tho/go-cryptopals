package challenge2

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func Example() {
	const (
		hexStr1 = "1c0111001f010100061a024b53535009181c"
		hexStr2 = "686974207468652062756c6c277320657965"
	)

	decoded1, _ := hex.DecodeString(hexStr1)
	decoded2, _ := hex.DecodeString(hexStr2)
	xor, _ := ComputeFixedXOR(decoded1, decoded2)

	fmt.Printf("%x\n", xor)
	// Output: 746865206b696420646f6e277420706c6179
}

func TestComputeFixedXOR(t *testing.T) {
	tests := map[string]struct {
		s1      []byte
		s2      []byte
		want    []byte
		wantErr bool
	}{
		"valid": {
			s1:      []byte{0x00, 0x00},
			s2:      []byte{0xff, 0xff},
			want:    []byte{0xff, 0xff},
			wantErr: false,
		},
		"nil s1": {
			s1:      nil,
			s2:      []byte{},
			want:    []byte{},
			wantErr: false,
		},
		"nil s2": {
			s1:      []byte{},
			s2:      nil,
			want:    []byte{},
			wantErr: false,
		},
		"nil s1 s2": {
			s1:      nil,
			s2:      nil,
			want:    []byte{},
			wantErr: false,
		},
		"blank s1 s2": {
			s1:      []byte{},
			s2:      []byte{},
			want:    []byte{},
			wantErr: false,
		},
		"unequal length buffers": {
			s1:      []byte{0x00},
			s2:      []byte{0xff, 0xff},
			want:    nil,
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s1 := append(tt.s1[:0:0], tt.s1...)
			s2 := append(tt.s2[:0:0], tt.s2...)

			got, err := ComputeFixedXOR(s1, s2)
			if (err != nil) != tt.wantErr {
				t.Fatalf("computeFixedXOR(%v, %v) err: %v, want: %v", tt.s1, tt.s2, err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeFixedXOR(%v, %v) = %v, want: %v", tt.s1, tt.s2, got, tt.want)
			}
			if !reflect.DeepEqual(s1, tt.s1) {
				t.Errorf("computeFixedXOR(%v, %v) changed first input buffer to %v", tt.s1, tt.s2, s1)
			}
			if !reflect.DeepEqual(s2, tt.s2) {
				t.Errorf("computeFixedXOR(%v, %v) changed second input buffer to %v", tt.s1, tt.s2, s2)
			}
		})
	}
}
