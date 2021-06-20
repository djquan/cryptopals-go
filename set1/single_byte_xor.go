package set1

import (
	"encoding/hex"
	"unicode"
)

func singleByteXor(input string) (byte, string, error) {
	bs, err := hex.DecodeString(input)

	if err != nil {
		return 0, "", err
	}

	score := 0
	var decoderByte byte
	var decodedString string

	for i := 0; i < 256; i++ {
		xored := make([]byte, len(bs))

		for j, b := range bs {
			xored[j] = b ^ byte(i)
		}
		newScore := scoreBs(xored)

		if newScore > score {
			decoderByte = byte(i)
			decodedString = string(xored)
			score = newScore
		}
	}

	return decoderByte, decodedString, nil
}

func scoreBs(bytes []byte) int {
	result := 0

	for _, b := range bytes {
		result += frequencyMap[byte(unicode.ToUpper(rune(b)))]
	}

	return result
}

var frequencyMap = map[byte]int{
	'E': 21912,
	'T': 16587,
	'A': 14810,
	'O': 14003,
	'I': 13318,
	'N': 12666,
	'S': 11450,
	'R': 10977,
	'H': 10795,
	'D': 7874,
	'L': 7253,
	'U': 5246,
	'C': 4943,
	'M': 4761,
	'F': 4200,
	'Y': 3853,
	'W': 3819,
	'G': 3693,
	'P': 3316,
	'B': 2715,
	'V': 2019,
	'K': 1257,
	'X': 315,
	'Q': 205,
	'J': 188,
	'Z': 128,
	' ': 16000,
}
