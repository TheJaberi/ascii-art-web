package asciiart

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// Color returns the ANSI escape sequence for the specified color
func ColorSelector(color string) string {
	var colors = map[string]string{
		"reset":      "\033[0m",
		"red":        "\033[31m",
		"green":      "\033[32m",
		"yellow":     "\033[33m",
		"blue":       "\033[34m",
		"purple":     "\033[35m",
		"cyan":       "\033[36m",
		"white":      "\033[37m",
		"gray":       "\033[90m",
		"darkred":    "\033[91m",
		"orange":     "\033[38;5;208m",
		"pink":       "\033[38;5;206m",
		"gold":       "\033[38;5;220m",
		"teal":       "\033[38;5;51m",
		"lavender":   "\033[38;5;183m",
		"brown":      "\033[38;5;130m",
		"lightblue":  "\033[38;5;39m",
		"magenta":    "\033[38;5;200m",
		"olive":      "\033[38;5;100m",
		"salmon":     "\033[38;5;203m",
		"skyblue":    "\033[38;5;111m",
		"darkpurple": "\033[38;5;53m",
		"lime":       "\033[38;5;46m",
		"maram":      "\033[38;5;206m",

		// Add more colors here as needed
	}
	if color1, ok := colors[color]; ok {
		return color1
	}

	// Check if the color is in RGB format (e.g., rgb(255, 0, 0))
	rgbRegex := regexp.MustCompile(`^rgb\((\d{1,3}),\s?(\d{1,3}),\s?(\d{1,3})\)$`)
	rgbMatches := rgbRegex.FindStringSubmatch(color)
	if len(rgbMatches) == 4 {
		r, _ := strconv.Atoi(rgbMatches[1])
		g, _ := strconv.Atoi(rgbMatches[2])
		b, _ := strconv.Atoi(rgbMatches[3])
		return RGBColor(r, g, b)
	}

	// Check if the color is in RGB hex format (#rrggbb)
	if isHexColor(color) {
		// Convert the hex color to RGB
		r, g, b := hexToRGB(color)
		return RGBColor(r, g, b)
	}

	// Check if the color is in HSL format (hsl(h, s, l))
	if isHSLColor(color) {
		// Convert the HSL color to RGB
		r, g, b := hslToRGB(color)
		return RGBColor(r, g, b)
	}

	// Print an error message for unrecognized colors
	fmt.Println("Unrecognized color:", color)
	return "" // Return an empty string for unrecognized colors
}

// rgbToRGB converts an RGB color to RGB values
func rgbToRGB(color string) (int, int, int) {
	// Extract the RGB values from the color string
	rgbPattern := regexp.MustCompile(`(\d+),\s*(\d+),\s*(\d+)`)
	matches := rgbPattern.FindStringSubmatch(color)
	if len(matches) != 4 {
		return 0, 0, 0
	}

	// Convert the RGB values to integers
	r, _ := strconv.Atoi(matches[1])
	g, _ := strconv.Atoi(matches[2])
	b, _ := strconv.Atoi(matches[3])

	return r, g, b
}

// RGBColor returns the ANSI escape sequence for an RGB color
func RGBColor(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// isHexColor checks if the color is in RGB hex format (#rrggbb)
func isHexColor(color string) bool {
	// Regular expression to match RGB hex format (#rrggbb)
	match, _ := regexp.MatchString(`^#([a-fA-F0-9]{6})$`, color)
	return match
}

// hexToRGB converts a hex color to RGB values
func hexToRGB(color string) (int, int, int) {
	// Remove the "#" symbol from the color
	color = color[1:]

	// Convert the color components to decimal values
	r, _ := strconv.ParseInt(color[:2], 16, 0)
	g, _ := strconv.ParseInt(color[2:4], 16, 0)
	b, _ := strconv.ParseInt(color[4:6], 16, 0)

	return int(r), int(g), int(b)
}

// isHSLColor checks if the color is in HSL format (hsl(h, s, l))
func isHSLColor(color string) bool {
	// Regular expression to match HSL format (hsl(h, s, l))
	match, _ := regexp.MatchString(`^hsl\(\d+,\s*\d+%?,\s*\d+%?\)$`, color)
	return match
}

// hslToRGB converts an HSL color to RGB values
func hslToRGB(color string) (int, int, int) {
	// Extract the HSL values from the color string
	hslPattern := regexp.MustCompile(`(\d+),\s*(\d+)%?,\s*(\d+)%?`)
	matches := hslPattern.FindStringSubmatch(color)
	if len(matches) != 4 {
		return 0, 0, 0
	}

	// Convert the HSL values to integers
	h, _ := strconv.Atoi(matches[1])
	s, _ := strconv.Atoi(matches[2])
	l, _ := strconv.Atoi(matches[3])

	// Normalize the HSL values
	h = h % 360
	s = int(math.Max(0, math.Min(100, float64(s))))
	l = int(math.Max(0, math.Min(100, float64(l))))

	// Convert HSL to RGB
	c := (1 - math.Abs(float64(2*l-100))/100) * float64(s) / 100
	x := c * (1 - math.Abs(float64((h/60)%2-1)))
	m := float64(l)/100 - c/2

	var r, g, b float64
	switch {
	case h < 60:
		r, g, b = c, x, 0
	case h < 120:
		r, g, b = x, c, 0
	case h < 180:
		r, g, b = 0, c, x
	case h < 240:
		r, g, b = 0, x, c
	case h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}

	// Scale the RGB values and convert to integers
	r = (r + m) * 255
	g = (g + m) * 255
	b = (b + m) * 255

	return int(r), int(g), int(b)
}
