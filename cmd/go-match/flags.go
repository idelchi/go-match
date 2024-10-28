package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// flags defines the command-line flags for the application.
func flags() {
	pflag.StringArray("include", []string{"**/*"}, "List of patterns to include")
	pflag.StringArray("exclude", nil, "List of patterns to exclude")
	pflag.Bool("version", false, "Show the version information and exit")
	pflag.BoolP("help", "h", false, "Show the help information and exit")

	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] [paths...]\n\n", "go-match")
		fmt.Fprintf(os.Stderr, "Check if paths match include patterns and don't match exclude patterns.\n\n")
		fmt.Fprintf(os.Stderr, "Arguments:\n")
		fmt.Fprintf(os.Stderr, "  paths   The paths to check (optional if piped through stdin)\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		pflag.PrintDefaults()
	}
}

// parseFlags parses the application configuration from command-line flags and environment variables.
func parseFlags() (cfg Config, err error) {
	flags()

	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return cfg, fmt.Errorf("binding flags: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("unmarshalling config: %w", err)
	}

	handleExitFlags(cfg)

	if err := validateInput(&cfg); err != nil {
		return cfg, fmt.Errorf("validating input: %w", err)
	}

	return cfg, nil
}

// ErrUsage represents an error that occurs due to incorrect usage of the application.
var ErrUsage = errors.New("usage error")

// validateInput validates the input provided to the application.
func validateInput(cfg *Config) error {
	cfg.Paths = pflag.Args()

	if len(cfg.Paths) == 0 {
		return fmt.Errorf("%w: no paths provided", ErrUsage)
	}

	// If both include and exclude are empty, fail
	if len(cfg.Include) == 0 && len(cfg.Exclude) == 0 {
		return fmt.Errorf("%w: no include or exclude patterns provided", ErrUsage)
	}

	return nil
}

//nolint:forbidigo // Function will print & exit for various help messages.
func handleExitFlags(cfg Config) {
	if viper.GetBool("version") {
		fmt.Println(version)
		os.Exit(0)
	}

	if viper.GetBool("help") {
		pflag.Usage()
		os.Exit(0)
	}

	if viper.GetBool("show") {
		fmt.Println(PrintJSON(cfg))

		os.Exit(0)
	}
}

// PrintJSON returns a pretty-printed JSON representation of the provided object.
func PrintJSON(obj any) string {
	bytes, err := json.MarshalIndent(obj, "  ", "    ")
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}
