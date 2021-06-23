package set1

import (
	"crypto/aes"
)

func decryptAESinECB(filename string, key []byte) ([]byte, error) {
	file, err := base64DecodeFile(filename)
	if err != nil {
		return nil, err
	}

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(file))

	for i := 0; i < len(file); i += len(key) {
		end := i + len(key)
		cipher.Decrypt(result[i:end], file[i:end])
	}

	return result, nil
}
