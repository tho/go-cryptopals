package challenge1

import (
	"encoding/base64"
	"encoding/hex"
)

// ConvertHexToBase64String returns the base64 encoding of the hexadecimal
// string s.
//
// ConvertHexToBase64String expects that s contains only hexadecimal characters
// and that s has even length. If the input is malformed, the function returns
// the encoded bytes before the error.
func ConvertHexToBase64String(s string) (string, error) {
	decoded, err := hex.DecodeString(s)
	return base64.StdEncoding.EncodeToString(decoded), err
}
