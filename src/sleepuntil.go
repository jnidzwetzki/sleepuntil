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

func tryDateFormats(userTimeValue string, dateFormats []string) (*time.Time, error) {
	var localLocation = time.Now().Local().Location()

	for _, layout := range dateFormats {
		t, err := time.ParseInLocation(layout, userTimeValue, localLocation)

		if err == nil {
			return &t, nil
		}
	}

	return nil, errors.New("Unable to parse date: " + userTimeValue)
}

func parseDate(userTimeValue string) (*time.Time, error) {
	var layouts = []string{"2006-01-02 15:04:05", "2006-01-02 15:04", "2006-01-02"}

	parsedValue, err := tryDateFormats(userTimeValue, layouts)

	if err == nil {
		return parsedValue, nil
	}

	// Try parsing with date string
	var currentTime = time.Now().Local()
	var datePrefix = currentTime.Format("2006-01-02")

	var userTimeValueWithPrefix = datePrefix + " " + userTimeValue

	parsedValue, err = tryDateFormats(userTimeValueWithPrefix, layouts)

	// Unable to parse, exit
	if err != nil {
		return nil, err
	}

	// Got hour value in the past (assume tomorrow is meant)
	if parsedValue.Before(currentTime) {
		var tomorrow = currentTime.AddDate(0, 0, 1)
		var datePrefix = tomorrow.Format("2006-01-02")
		var userTimeValueWithPrefixTomorrow = datePrefix + " " + userTimeValue
		return tryDateFormats(userTimeValueWithPrefixTomorrow, layouts)
	}

	return parsedValue, nil
}

func showAnimation() {

}

func main() {
	flagErr := globalFlagSet.Parse(os.Args[1:])

	if flagErr != nil {
		showHelpAndExit()
	}

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

	go showAnimation()

	var currentTime = time.Now().Local()
	var sleepTime = parseDate.Sub(currentTime)

	fmt.Printf("Sleep time %f\n", sleepTime.Seconds())

	time.Sleep(sleepTime)
}
