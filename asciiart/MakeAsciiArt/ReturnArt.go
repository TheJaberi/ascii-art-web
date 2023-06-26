package asciiart

import (
	"fmt"
	"strings"
)

func ReturnArt(str string, subStr string, option []string,fileName string, width int) string{
	asciiArr, err := AsciiArtArray(fileName)
	if err != nil {
		fmt.Println("Error", err)
		return "Error"
	}
	printArr := MakeTheArt(str, subStr, asciiArr, option, width)
	// align
	for i, asciiline := range printArr {
		if GetLength(asciiline) <= width || option[2] == "" {
			printArr[i] = AlignString(asciiline, width, option[2])
		} else {
			fmt.Println("Error: The text does not fit in terminal")
			break
		}
	}
	asciiString := strings.Join(printArr,"\n")
	return asciiString
}