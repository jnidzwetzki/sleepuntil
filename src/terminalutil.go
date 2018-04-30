package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

// Result is: 30 130
func determineTerminalSizeRaw() (string, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	var commandOutput = string(out)

	return commandOutput, err
}

// Parse: 30 130
func parseTerminalSizeRaw(sizeRaw string) (int, int, error) {
	regex, err := regexp.Compile(`(\d+) (\d+)`)

	if err != nil {
		log.Fatalf("Got exception while compiling regex %s\n", err)
		os.Exit(-1)
	}

	result := regex.FindStringSubmatch(sizeRaw)

	if len(result) != 3 {
		return -1, -1, fmt.Errorf("The array has not three elements: %s", result)
	}

	height, err := strconv.Atoi(result[1])

	if err != nil {
		return -1, -1, err
	}

	width, err := strconv.Atoi(result[2])

	if err != nil {
		return -1, -1, err
	}

	return width, height, nil
}

// GetWidth returns the width of the terminal
func getTeminalWidth() (int, error) {
	commandOutput, err := determineTerminalSizeRaw()

	if err != nil {
		return 0, err
	}

	width, _, err := parseTerminalSizeRaw(commandOutput)

	return width, err
}
