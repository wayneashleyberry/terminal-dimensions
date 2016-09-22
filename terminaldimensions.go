package terminaldimensions

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func parts() []string {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	return strings.Split(string(out), " ")
}

// Width gives you the width of the terminal
func Width() (uint, error) {
	parts := parts()
	x, err := strconv.Atoi(strings.Replace(parts[1], "\n", "", -1))
	if err != nil {
		return 0, err
	}
	return uint(x), err
}

// Height gives you the height of the terminal
func Height() (uint, error) {
	parts := parts()
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	return uint(x), err
}
