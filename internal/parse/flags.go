// Package parse handles configuration parsing from command-line flags
// and environment variables, with input validation and format detection.
package parse

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/idelchi/go-match/internal/commands"
	"github.com/idelchi/go-match/internal/config"
	"github.com/idelchi/godyl/pkg/pretty"
)

// Parse parses the application configuration (in order of precedence) from:
//   - command-line flags
//   - environment variables
func Parse(version string) (cfg config.Config, err error) {
	commands.Flags()

	// Parse the command-line flags
	pflag.Parse()

	// Bind pflag flags to viper
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return cfg, fmt.Errorf("binding flags: %w", err)
	}

	// Set viper to automatically read from environment variables
	viper.SetEnvPrefix("go_match")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Unmarshal the configuration into the Config struct
	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("unmarshalling config: %w", err)
	}

	// Validate the input
	err = validateInput(&cfg)

	// Handle the commandline flags that exit the application
	handleExitFlags(cfg, version)

	if err != nil {
		return cfg, fmt.Errorf("validating input: %w", err)
	}

	return cfg, nil
}

//nolint:forbidigo // Function will print & exit for various help messages.
func handleExitFlags(cfg config.Config, version string) {
	// Check if the version flag was provided
	if viper.GetBool("version") {
		fmt.Println(version)

		os.Exit(0)
	}

	// Check if the help flag was provided
	if viper.GetBool("help") {
		pflag.Usage()

		os.Exit(0)
	}

	if viper.GetBool("show") {
		pretty.PrintJSON(cfg)

		os.Exit(0)
	}
}

// validateInput validates the input provided to the application.
func validateInput(cfg *config.Config) error {
	cfg.Paths = pflag.Args()

	if len(cfg.Paths) == 0 {
		return fmt.Errorf("%w: no paths provided", config.ErrUsage)
	}

	// If both include and exclude are empty, fail
	if len(cfg.Include) == 0 && len(cfg.Exclude) == 0 {
		return fmt.Errorf("%w: no include or exclude patterns provided", config.ErrUsage)
	}

	return nil
}
