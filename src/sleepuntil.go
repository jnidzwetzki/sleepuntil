package main

import (
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

	if globalFlags.Animate {
		go showAnimation()
	}

	var currentTime = time.Now().Local()
	var sleepTime = parseDate.Sub(currentTime)

	if globalFlags.Verbose {
		fmt.Printf("Sleep time %f\n", sleepTime.Seconds())
	}

	time.Sleep(sleepTime)
}
