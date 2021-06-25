package set1

import (
	"bufio"
	"encoding/hex"
	"errors"
	"os"
)

func detectAES(filename string) ([]byte, error) {
	input, err := readAndHexDecodeFile(filename)
	if err != nil {
		return nil, err
	}

	for _, v := range input {
		seen := make(map[string]bool)

		for i := 0; i < len(v); i += 16 {
			end := i + 16
			chunk := v[i:end]

			if _, ok := seen[string(chunk)]; ok {
				result := make([]byte, hex.EncodedLen(len(v)))
				hex.Encode(result, v)
				return result, nil
			}

			seen[string(chunk)] = true
		}

	}
	return nil, errors.New("didn't find an AES")
}

func readAndHexDecodeFile(filename string) ([][]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	input := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		bytes := []byte(scanner.Text())
		decoded := make([]byte, hex.DecodedLen(len(bytes)))
		_, err := hex.Decode(decoded, bytes)
		if err != nil {
			return nil, err
		}

		input = append(input, decoded)
	}

	return input, nil
}
