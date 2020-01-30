package challenge1

import (
	"fmt"
	"testing"
)

func Example() {
	const hexStr = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	base64Str, _ := ConvertHexToBase64String(hexStr)

	fmt.Println(base64Str)
	// Output: SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
}

func TestConvertHexToBase64String(t *testing.T) {
	tests := map[string]struct {
		s       string
		want    string
		wantErr bool
	}{
		"empty": {
			s:       "",
			want:    "",
			wantErr: false,
		},
		"valid": {
			s:       "0123456789abcdef",
			want:    "ASNFZ4mrze8=",
			wantErr: false,
		},
		"partially valid": {
			s:       "0123456789abcdefghijkl",
			want:    "ASNFZ4mrze8=",
			wantErr: true,
		},
		"invalid": {
			s:       "ghijkl",
			want:    "",
			wantErr: true,
		},
		"partially invalid": {
			s:       "ghijkl0123456789abcdef",
			want:    "",
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ConvertHexToBase64String(tt.s)
			if (err != nil) != tt.wantErr {
				t.Fatalf("convertHexToBase64String(%q) err: %v, want: %t", tt.s, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("convertHexToBase64String(%q) = %q, want: %q", tt.s, got, tt.want)
			}
		})
	}
}
