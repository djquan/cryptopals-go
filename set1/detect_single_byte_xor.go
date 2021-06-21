package set1

import (
	"bufio"
	"os"
)

func detectSingleByteXor(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var message string
	highScore := 0
	for scanner.Scan() {
		_, s, score, err := singleByteXor(scanner.Text())
		if err != nil {
			return "", err
		}

		if score > highScore {
			highScore = score
			message = s
		}

	}

	return message, nil
}
