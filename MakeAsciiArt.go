package asciiart

import (
	"strings"
)

func MakeTheArt(str string, subStr string, asciiArr []string, option []string, width int) []string {
	// Split the input text into words and check if it contains only new lines
	words := strings.Split(str, "\\n")
	if AllEmpty(words) {
		words = words[1:]
	}
	printArr := []string{}
	// Iterate over each word in the input text
	for _, currentWord := range words {
		// Skip empty words and print a newline
		if currentWord == "" {
			printArr = append(printArr, "")
			continue
		}
		if option[2] == "justify" {
			// Trim leading and trailing spaces
			currentWord = strings.TrimSpace(currentWord)
			// Replace extra spaces within the string with a single space
			currentWord = strings.Join(strings.Fields(currentWord), " ")
		}
		printArr = MakeString(currentWord, subStr, asciiArr, printArr, option, width)
	}
	return printArr
}

func MakeString(str string, subStr string, asciiArr []string, printArr []string, option []string, width int) []string {
	for i := 0; i < 8; i++ {
		line := MakeAsciiArtLine(str, subStr, asciiArr, option, i, width)
		printArr = append(printArr, line)
	}
	return printArr
}

func MakeAsciiArtLine(str string, subStr string, asciiArr []string, option []string, i int, width int) string {
	color := option[1]
	currentLine := ""
	spaces := "      " // spaces in the file
	if option[2] == "justify" && len(strings.Split(str, " ")) > 1 {
		numOfSpaces := CalcJustifySpaces(str, asciiArr, i, width)
		spaces = strings.Repeat(" ", numOfSpaces)
	}
	for pos, char := range str {
		if char == ' ' {
			currentLine += spaces
		} else if ToColor(str, subStr, pos) && option[0] == "" && option[1] != "" {
			currentLine += color + asciiArr[((int(char)-32)*8)+i] + "\x1b[0m"
		} else {
			currentLine += asciiArr[((int(char)-32)*8)+i]
		}
	}
	return currentLine
}
