package AMJ

import (
	"bufio"
	"errors"
	"os"
)

func GetLine(num int, filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil { // Checking if we have the right file name
		return "", errors.New("500: Internal Server Error: Error opening file: " + err.Error())
	}
	defer file.Close()

	str := ""
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		if lineNumber == num {
			str = scanner.Text()
		}
		lineNumber++
	}
	return str, nil
}
