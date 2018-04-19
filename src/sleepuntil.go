package main

import (
	"flag"
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
		Time    string
	}{}
)

func init() {
	out = new(tabwriter.Writer)
	out.Init(os.Stdout, 0, 8, 1, '\t', 0)
	globalFlagSet.BoolVar(&globalFlags.Verbose, "verbose", false, "Be verbose.")
	globalFlagSet.StringVar(&globalFlags.Time, "time", "12:00", "The time to wait for")
}

func main() {
	globalFlagSet.Parse(os.Args[1:])
	var args = globalFlagSet.Args()

	if len(args) < 1 {
		args = append(args, "help")
	}
}
