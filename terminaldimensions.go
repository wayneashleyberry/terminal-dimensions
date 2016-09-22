package terminaldimensions

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Width gives you the width of the terminal
func Width() (uint, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, outputErr := cmd.Output()
	if outputErr != nil {
		return 0, outputErr
	}
	parts := strings.Split(string(out), " ")
	x, err := strconv.Atoi(strings.Replace(parts[1], "\n", "", -1))
	if err != nil {
		return 0, err
	}
	return uint(x), err
}
