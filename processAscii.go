package asciiart

import (
	"fmt"
	"strings"
)

func ProcessAscii(str, subStr, font, color, alignment string) []string {
	var asciiLines []string

	fileName := font + ".txt"
	option := []string{"", "", ""} // Option[0] for Output , Option[1] for color , Option[2] for align

	// set the str and subStr values
	if subStr != "" {
		if len(str) > len(subStr) {
			str, subStr = subStr, str
		}
		if !strings.Contains(str, subStr) {
			fmt.Println("String does not contain the substring")
			return asciiLines
		}
	}

	// If there is an align option, the width of the terminal will be calculated
	width, err := CalcWidth()
	if err != nil {
		fmt.Println("Error:", err)
		return asciiLines
	}

	asciiArr, err := AsciiArtArray(fileName)
	if err != nil {
		fmt.Println("Error", err)
		return asciiLines
	}

	printArr := MakeTheArt(str, subStr, asciiArr, option, width)
	for _, asciiline := range printArr {
		if GetLength(asciiline) <= width || option[2] == "" {
			finalLine := AlignString(asciiline, width, option[2])
			asciiLines = append(asciiLines, finalLine)
		} else {
			fmt.Println("Error: The text does not fit in terminal")
			break
		}
	}

	return asciiLines
}
