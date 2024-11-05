// Package config defines and validates the configuration structure for version tag generation and formatting.
package config

import (
	"errors"
	"fmt"
)

// ErrUsage indicates an error in command-line usage or configuration.
var ErrUsage = errors.New("usage error")

// Config represents the configuration for the go-match application.
type Config struct {
	// Include is a list of patterns to include.
	Include []string

	// Exclude is a list of patterns to exclude.
	Exclude []string

	// Paths is a list of paths to check.
	Paths []string
}

// Validate validates the configuration.
func (c Config) Validate() error {
	if len(c.Include) == 0 && len(c.Exclude) == 0 {
		return fmt.Errorf("%w: at least one include or exclude pattern must be provided", ErrUsage)
	}

	if len(c.Paths) == 0 {
		return fmt.Errorf("%w: at least one path must be provided", ErrUsage)
	}

	return nil
}
