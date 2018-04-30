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

// Print a wait animation
func showAnimation(startTime time.Time, sleepTime time.Duration) {
	width, err := getTeminalWidth()

	if err != nil {
		log.Fatal("Got error while running animation")
	}

	for {
		var currentTime = time.Now().Local()
		elapsedTime := currentTime.Sub(startTime)
		percent := sleepTime.Seconds() / elapsedTime.Seconds()

		fmt.Print("\r")
		fmt.Print("|")

		elementsToPrint := width - 10

		for i := 0; i < elementsToPrint; i++ {
			if (float64(elementsToPrint) / float64(i)) < percent {
				fmt.Print("-")
			} else {
				fmt.Print("*")
			}
		}

		fmt.Print("|")
		time.Sleep(1 * time.Second)
	}

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

	var currentTime = time.Now().Local()
	var sleepTime = parseDate.Sub(currentTime)

	if globalFlags.Animate {
		go showAnimation(currentTime, sleepTime)
	}

	if globalFlags.Verbose {
		fmt.Printf("Sleep time %f\n", sleepTime.Seconds())
	}

	time.Sleep(sleepTime)

	fmt.Println("")
}
