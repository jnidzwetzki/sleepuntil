package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
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
		Help    bool
	}{}
)

func init() {
	globalFlagSet.BoolVar(&globalFlags.Verbose, "verbose", false, "Be verbose.")
	globalFlagSet.BoolVar(&globalFlags.Help, "help", false, "Show help.")
	globalFlagSet.BoolVar(&globalFlags.Animate, "progress", false, "Show a progress animation for the remaining time.")
}

func showHelpAndExit() {
	fmt.Printf("Usage: %s <time> <flags>\n\n", cliName)

	globalFlagSet.PrintDefaults()

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
	width, err := getTeminalWidth()

	if err != nil {
		log.Fatal("Got error while running animation")
	}

	fmt.Printf("Terminal width is: %d\n", width)
}

func main() {
	flagErr := globalFlagSet.Parse(os.Args[2:])

	if flagErr != nil || globalFlags.Help {
		showHelpAndExit()
	}

	if len(os.Args) < 1 {
		showHelpAndExit()
	}

	dateToParse := os.Args[1]
	parseDate, err := parseDate(dateToParse)

	if err != nil {
		fmt.Printf("Unable to parse %s, exiting\n", dateToParse)
		os.Exit(-1)
	}

	fmt.Printf("Wait until %s\n", *parseDate)

	if globalFlags.Animate {
		go showAnimation()
	}

	var currentTime = time.Now().Local()
	var sleepTime = parseDate.Sub(currentTime)

	fmt.Printf("Sleep time %f\n", sleepTime.Seconds())

	time.Sleep(sleepTime)
}
