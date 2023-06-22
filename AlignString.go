package asciiart

import (
	"fmt"
	"regexp"
	"strings"
)

func AlignString(str string, width int, align string) string {
	// Find the length of the string without the color
	strLen := GetLength(str)
	if align == "center" {
		// Calculate the number of spaces needed on each side of the string
		spaces := width - strLen
		leftSpaces := spaces / 2
		rightSpaces := spaces / 2
		// Use Sprintf to format the string with the correct number of spaces
		centeredString := fmt.Sprintf("%s%s%s", strings.Repeat(" ", leftSpaces), str, strings.Repeat(" ", rightSpaces))
		return centeredString
	} else if align == "left" {
		// Use Sprintf to format the string with the correct number of spaces
		leftString := fmt.Sprintf("%s%s", str, strings.Repeat(" ", width-strLen))
		return leftString
	} else if align == "right" {
		rightString := fmt.Sprintf("%s%s", strings.Repeat(" ", width-strLen), str)
		return rightString
	}
	return str
}

func GetLength(str string) int {
	// remove the color
	re := regexp.MustCompile(`\x1b\[[0-9;]*[mK]`)
	str = re.ReplaceAllString(str, "")
	return len(str)
}

func CalcJustifySpaces(str string, asciiArr []string, i int, width int) int {
	words := strings.Fields(str)
	// the string without the spaces
	currentLine := ""
	for _, char := range str {
		if char != ' ' {
			currentLine += asciiArr[((int(char)-32)*8)+i]
		}
	}
	spacesBet := width - len(currentLine)
	spacesBet = spacesBet / (len(words) - 1)
	return spacesBet
}
