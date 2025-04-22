// Package logic implements the core logic of the application.
package logic

import (
	"encoding/json"
	"fmt"

	"github.com/idelchi/go-match/internal/match"
	"github.com/idelchi/go-match/internal/parse"
)

// Run executes the core logic of the application.
func Run(version string) error {
	cfg, err := parse.Parse(version)
	if err != nil {
		return fmt.Errorf("parsing flags: %w", err)
	}

	if err = cfg.Validate(); err != nil {
		return fmt.Errorf("application configuration: %w", err)
	}

	matches, err := match.Match(cfg.Paths, cfg.Include, cfg.Exclude)
	if err != nil {
		return fmt.Errorf("matching paths: %w", err)
	}

	jsonOutput, err := json.Marshal(matches)
	if err != nil {
		return fmt.Errorf("marshalling JSON: %w", err)
	}

	fmt.Println(string(jsonOutput)) //nolint:forbidigo // Print the JSON output

	return nil
}
