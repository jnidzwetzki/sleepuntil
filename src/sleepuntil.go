package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	cliVersion    = 0.1
	cliName       = "sleepuntil"
	cliDesription = "Sleep until a certain time is reached"
)

var (
	globalFlagSet = flag.NewFlagSet("sleepuntil", flag.ExitOnError)

	globalFlags = struct {
		Verbose bool
		Animate bool
		Time    string
	}{}
)

func init() {
	globalFlagSet.BoolVar(&globalFlags.Verbose, "verbose", false, "Be verbose.")
	globalFlagSet.BoolVar(&globalFlags.Animate, "animate", false, "Show an animation for the remaining time")
}

func showHelpAndExit() {
	fmt.Printf("Usage: %s <time> <flags>\n\n", cliName)

	myFlags := make([]*flag.Flag, 0)
	globalFlagSet.VisitAll(func(f *flag.Flag) {
		myFlags = append(myFlags, f)
	})

	fmt.Printf("Supported flags:\n")
	for _, flag := range myFlags {
		fmt.Printf("-%s (Default: %s) %s\n", flag.Name, flag.Value, flag.Usage)
	}

	fmt.Printf("\n")
	fmt.Printf("Version: %g\n", cliVersion)
	fmt.Printf("Please report bugs here: https://github.com/jnidzwetzki/sleepuntil/issues\n")
	fmt.Printf("\n")

	os.Exit(-1)
}

func parseDate(dateToParse string) (*time.Time, error) {
	dateLayouts := []string{"15:04", "15:04:05", "2006-01-02T15:04:05.000Z"}

	for _, layout := range dateLayouts {
		t, err := time.Parse(layout, dateToParse)

		if err == nil {
			return &t, nil
		}
	}

	return nil, errors.New("Unable to parse date: " + dateToParse)
}

func main() {
	globalFlagSet.Parse(os.Args[1:])
	var args = globalFlagSet.Args()

	if len(args) < 1 {
		showHelpAndExit()
	}

	parseDate, err := parseDate(args[0])

	if err != nil {
		fmt.Printf("Unable to parse %s, exiting\n", args[0])
		os.Exit(-1)
	}

	fmt.Printf("Wait until %s\n", *parseDate)
}
