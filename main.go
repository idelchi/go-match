// go-match is a simple command-line utility to check if a path matches a globstar pattern.
package main

import (
	"fmt"
	"os"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/spf13/pflag"
)

func main() {
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] <path> <pattern>\n\n", "go-match")
		fmt.Fprintf(os.Stderr, "Check if a path matches a globstar pattern.\n\n")
		fmt.Fprintf(os.Stderr, "Arguments:\n")
		fmt.Fprintf(os.Stderr, "  path    The path to check\n")
		fmt.Fprintf(os.Stderr, "  pattern The globstar pattern to match against\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		pflag.PrintDefaults()
	}

	pflag.Parse()

	const expectedNumArgs = 2

	if pflag.NArg() != expectedNumArgs {
		pflag.Usage()

		os.Exit(1)
	}

	path := pflag.Arg(0)
	pattern := pflag.Arg(1)

	match, err := doublestar.Match(pattern, path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	if match {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	os.Exit(0)
}
