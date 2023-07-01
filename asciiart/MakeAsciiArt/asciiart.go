package AMJ

import (
	"errors"
	"strings"
)

func Asciiart(str, font string) (string, error) { // used for ascii art web
	filename, err := Banners(font)
	if err != nil {
		return "", err
	}
	restoprint := "" // for the output file

	// replacing \r\n with \\n so our ascii art program can deal with it
	strarr := strings.Split(str, "\r\n")
	str = strings.Join(strarr, "\\n")
	for _, r := range str {
		if r < 32 || r > 126 {
			return "", errors.New("400: Bad Request: Use English chars only")
		}
	}
	str = Alphabetformat(str) // fixing the \t
	res := ""
	args := strings.Split(str, "\\n")
	if AllEmpty(args) {
		args = args[1:]
	}
	// Writing text line by line into res
	for _, word := range args {
		if word == "" {
			// fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {
			res, err = WhatToPrint(i, word, filename, res)
			if err != nil {
				return "", err
			}
			// fmt.Println(res)

			restoprint = restoprint + "\n" + res
			res = ""
		}
	}

	return restoprint[1:], nil
}
