package set1

import (
	"encoding/base64"
	"encoding/hex"
)

func hexToBase64(input string) ([]byte, error) {
	in, err := hex.DecodeString(input)

	if err != nil {
		return []byte{}, err
	}

	encoder := base64.StdEncoding
	result := make([]byte, encoder.EncodedLen(len(in)))
	encoder.Encode(result, in)

	return result, nil
}
