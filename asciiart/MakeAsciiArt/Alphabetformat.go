package AMJ

import (
	"strings"
)

func Alphabetformat(s string) string {
	s=strings.ReplaceAll(s,"\\t","    ")
	return s
}
