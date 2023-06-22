package asciiart

import (
	"fmt"
	"strings"
)

func printArray(arr []string) {
	// Loop through the array and print each element
	for i, element := range arr {
		fmt.Println(element)

		// If the current index is a multiple of the interval, print the separator
		if (i+1)%8 == 0 && i != len(arr)-1 {
			fmt.Println(strings.Repeat("-", 20))
		}
	}
}
