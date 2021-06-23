package set1

import "encoding/hex"

func fixedXor(s1, s2 string) ([]byte, error) {
	result := make([]byte, len(s1))

	b1, err := hex.DecodeString(s1)
	if err != nil {
		return nil, err
	}

	b2, err := hex.DecodeString(s2)
	if err != nil {
		return nil, err
	}

	xorB := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		xorB[i] = b1[i] ^ b2[i]
	}

	hex.Encode(result, xorB)
	return result, nil
}
