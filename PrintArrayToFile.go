package asciiart

import (
	"fmt"
	"os"
	"regexp"
)

func PrintArrayToFile(array []string, width int, option []string) error {
	file, err := os.Create(option[0])
	if err != nil {
		return err
	}
	defer file.Close()
	for _, str := range array {
		str = AlignString(str, width, option[2])
		// remove the color
		re := regexp.MustCompile(`\x1b\[[0-9;]*[mK]`)
		str = re.ReplaceAllString(str, "")
		fmt.Fprintln(file, str)
	}
	fmt.Fprintln(file, "")

	return nil
}
