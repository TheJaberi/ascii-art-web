package asciiart

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func CalcWidth() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return -1, err
	}

	// Parse the output of the command to get the width
	size := strings.Split(string(out), " ")
	width, err := strconv.Atoi(strings.TrimSpace(size[1]))
	if err != nil {
		return -1, err
	}
	return width, nil
}
