package set1

import (
	"encoding/hex"
)

func repeatingKeyXor(plaintext, key string) []byte {
	bs := []byte(plaintext)
	ciphertextBytes := make([]byte, len(bs))
	keyBs := []byte(key)

	for i, b := range bs {
		keyPart := keyBs[i%len(keyBs)]

		ciphertextBytes[i] = b ^ keyPart
	}

	r := make([]byte, hex.EncodedLen(len(ciphertextBytes)))
	hex.Encode(r, ciphertextBytes)

	return r
}
