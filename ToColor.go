package asciiart

import (
	"strings"
)

func ToColor(str string, substr string, myPos int) bool {
	if substr == "" {
		return true
	}
	AllPos := FindAllSubstringPos(str, substr) // Array that contains all positions of substrings
	if AllPos == nil {                         // str does not contain the substring
		return false
	}
	for _, aPos := range AllPos {
		if myPos >= aPos && myPos < aPos+len(substr) {
			return true
		}
	}
	return false
}

func FindAllSubstringPos(str string, substr string) []int {
	var SubPos []int
	lastIndex := 0
	for {
		firstPos := strings.Index(str[lastIndex:], substr)
		if firstPos == -1 || substr == "" { // str[lastIndex:] does not contain the substring
			break
		}
		firstPos += lastIndex
		SubPos = append(SubPos, firstPos)
		lastIndex = firstPos + len(substr)
		if lastIndex >= len(str) { // we reached the end of the string
			break
		}
	}
	return SubPos
}
