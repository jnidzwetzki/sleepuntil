package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

const (
	cliName       = "sleepuntil"
	cliDesription = "Sleep until a certain time is reached"
)

var (
	out *tabwriter.Writer

	globalFlagSet = flag.NewFlagSet("sleepuntil", flag.ExitOnError)

	globalFlags = struct {
		Verbose bool
		Animate bool
		Time    string
	}{}
)

func init() {
	out = new(tabwriter.Writer)
	out.Init(os.Stdout, 0, 8, 1, '\t', 0)
	globalFlagSet.BoolVar(&globalFlags.Verbose, "verbose", false, "Be verbose.")
	globalFlagSet.BoolVar(&globalFlags.Animate, "animate", false, "Show an animation for the remaining time")
}

func showHelp() {
	fmt.Printf("Usage: %s <time> <parameter>\n", cliName)

	myFlags := make([]*flag.Flag, 0)
	globalFlagSet.VisitAll(func(f *flag.Flag) {
		myFlags = append(myFlags, f)
	})

	for _, flag := range myFlags {
		fmt.Printf("Flag: -%s (Default: %s) %s\n", flag.Name, flag.Value, flag.Usage)
	}
}

func main() {
	globalFlagSet.Parse(os.Args[1:])
	var args = globalFlagSet.Args()

	if len(args) < 1 {
		showHelp()
	}
}
