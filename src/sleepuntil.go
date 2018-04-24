package main

import (
	"flag"
	"fmt"
	"os"
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

func showHelp() {
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
}

func main() {
	globalFlagSet.Parse(os.Args[1:])
	var args = globalFlagSet.Args()

	if len(args) < 1 {
		showHelp()
	}
}
