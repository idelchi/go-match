// Package commands provides command-line flag configuration
// and parsing for determining version tag formats and bump rules.
package commands

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

// Flags defines the command-line flags for the application.
func Flags() {
	pflag.BoolP("help", "h", false, "Show the help information and exit")
	pflag.BoolP("version", "v", false, "Show the version information and exit")
	pflag.BoolP("show", "s", false, "Show the configuration and exit")

	pflag.StringArrayP("include", "i", []string{"**/*"}, "List of patterns to include")
	pflag.StringArrayP("exclude", "e", nil, "List of patterns to exclude")

	pflag.CommandLine.SortFlags = false

	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] [paths...]\n\n", "go-match")
		fmt.Fprintf(os.Stderr, "Check if paths match include patterns and don't match exclude patterns.\n\n")
		fmt.Fprintf(os.Stderr, "Arguments:\n")
		fmt.Fprintf(os.Stderr, "  paths   The paths to check\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		pflag.PrintDefaults()
	}
}
