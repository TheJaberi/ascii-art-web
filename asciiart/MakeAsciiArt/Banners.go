package AMJ

import (
	"errors"
	"strings"
)

func Banners(banner string) (string, error) {
	banners := map[string]string{
		"shadow":     "shadow.txt",
		"standard":   "standard.txt",
		"thinkertoy": "thinkertoy.txt",
		"doom":       "doom.txt",
		"extrafont":  "extrafont.txt",
		// Add more banners here as needed
	}

	banner = strings.ToLower(banner)

	if banner1, ok := banners[banner]; ok {
		style := "asciiart/Fonts/" + banner1
		return style, nil
	} else {
		return "", errors.New("500: Internal Server Error")
	}
}
