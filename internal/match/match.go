// Package match provides a function to match paths against include and exclude patterns.
package match

import (
	"fmt"

	"github.com/bmatcuk/doublestar/v4"
)

// Match returns the paths that match the include patterns and don't match the exclude patterns.
// The include and exclude patterns are matched using the globstar syntax.
// If the include patterns are empty, all paths are included.
// Exclude patterns always take precedence over include patterns.
func Match(paths, includes, excludes []string) ([]string, error) {
	matchedPaths := []string{}

	for _, path := range paths {
		included := len(includes) == 0

		for _, pattern := range includes {
			match, err := doublestar.Match(pattern, path)
			if err != nil {
				return nil, fmt.Errorf("matching include pattern %s: %w", pattern, err)
			}

			if match {
				included = true

				break
			}
		}

		excluded := false

		for _, pattern := range excludes {
			match, err := doublestar.Match(pattern, path)
			if err != nil {
				return nil, fmt.Errorf("matching exclude pattern %s: %w", pattern, err)
			}

			if match {
				excluded = true

				break
			}
		}

		if included && !excluded {
			matchedPaths = append(matchedPaths, path)
		}
	}

	return matchedPaths, nil
}
