package asciiart

import (
	"bufio"
	"fmt"
	"os"
)

func AsciiArtArray(fileName string) ([]string, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	// start scanning the file and append each line into the array
	scanner := bufio.NewScanner(file)
	var lines []string
	i := 0
	for scanner.Scan() {
		if i%9 != 0 {
			lines = append(lines, scanner.Text())
		}
		i++
	}
	return lines, nil
}
