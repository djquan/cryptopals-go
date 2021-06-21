package set1

import (
	"bufio"
	"encoding/base64"
	"os"
)

func breakTheCipher(filename string) ([]byte, error) {
	decodeFile, err := base64DecodeFile(filename)
	if err != nil {
		return []byte{}, err
	}

	size := findKeySizes(decodeFile)
	blocks := extractBlocks(decodeFile, size)
	transposed := transposeBlocks(size, decodeFile, blocks)
	result := xorBlocks(size, transposed)

	return result, nil
}

func xorBlocks(size int, transposed [][]byte) []byte {
	result := make([]byte, size)
	for i, bytes := range transposed {
		xor, _, _ := singleByteXorBytes(bytes)
		result[i] = xor
	}
	return result
}

func transposeBlocks(size int, decodeFile []byte, blocks [][]byte) [][]byte {
	transposed := make([][]byte, size)

	for i := 0; i < size; i++ {
		transposed[i] = make([]byte, len(decodeFile)/size+1)
		for j, block := range blocks {
			transposed[i][j] = block[i]
		}
	}
	return transposed
}

func extractBlocks(decodeFile []byte, size int) [][]byte {
	blocks := make([][]byte, len(decodeFile)/size)
	for i := range blocks {
		blocks[i] = make([]byte, size)
	}

	for i, b := range decodeFile {
		if i/size >= len(blocks) {
			break // partial block
		}
		blocks[i/size][i%size] = b
	}
	return blocks
}

func findKeySizes(input []byte) int {
	var likelyKeySize int
	var minHamming = -1

	for i := 2; i <= 40; i++ {
		distance := 0

		for j := 0; j < len(input)-i-1-i; j += i {
			a := input[j : i+j]
			b := input[i+j : i+j+i]
			distance += hammingDistance(a, b)
		}

		distance /= i
		distance /= len(input) / i

		if minHamming < 0 || minHamming > distance {
			minHamming = distance
			likelyKeySize = i
		}
	}

	return likelyKeySize
}

func base64DecodeFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []byte{}, err
	}

	input := make([]byte, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		input = append(input, []byte(scanner.Text())...)
	}

	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(input)))
	_, err = base64.StdEncoding.Decode(decoded, input)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func hammingDistance(a, b []byte) int {
	result := 0
	for i := 0; i < len(a); i++ {
		b1 := a[i]
		b2 := b[i]

		for i := 0; i < 8; i++ {
			mask := byte(1 << i)

			if (b1 & mask) != (b2 & mask) {
				result++
			}
		}
	}

	return result
}
