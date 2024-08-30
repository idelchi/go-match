package main

import (
	"fmt"
)

// Config represents the configuration for the go-match application.
type Config struct {
	Include []string
	Exclude []string
	Paths   []string
}

// Validate validates the configuration.
func (c Config) Validate() error {
	if len(c.Include) == 0 && len(c.Exclude) == 0 {
		return fmt.Errorf("%w: at least one include or exclude pattern must be provided", ErrUsage)
	}

	return nil
}
