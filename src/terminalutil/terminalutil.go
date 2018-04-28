package terminalutil

import (
	"os"
	"os/exec"
	"strings"
)

// Result is: 30 130
func determineSize() (string, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	var commandOutput = string(out)

	return commandOutput, err
}

// GetWidth returns the width of the terminal
func GetWidth() (int, error) {
	commandOutput, err := determineSize()

	if err != nil {
		return 0, err
	}

	strings.Split(commandOutput, " ")

	return 0, nil
}
