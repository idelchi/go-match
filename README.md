# go-match

[![Go Reference](https://pkg.go.dev/badge/github.com/idelchi/go-match.svg)](https://pkg.go.dev/github.com/idelchi/go-match)
[![Go Report Card](https://goreportcard.com/badge/github.com/idelchi/go-match)](https://goreportcard.com/report/github.com/idelchi/go-match)
[![Build Status](https://github.com/idelchi/go-match/actions/workflows/github-actions.yml/badge.svg)](https://github.com/idelchi/go-match/actions/workflows/github-actions.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`go-match` is a simple command-line utility that filters a list of paths based on include and exclude globstar patterns.

The program takes a list of paths as input and outputs a JSON array of paths that
match the include patterns and don't match the exclude patterns.

## Installation

```sh
curl -sSL https://raw.githubusercontent.com/idelchi/go-match/refs/heads/main/install.sh | sh -s -- -d ~/.local/bin
```

## Usage

```sh
go-match [flags] [paths...]
```

### Configuration

| Flag            | Environment Variable | Description                                   | Default |
| --------------- | -------------------- | --------------------------------------------- | ------- |
| `--include`     | `GO_MATCH_INCLUDE`   | Include patterns (can be used multiple times) | -       |
| `--exclude`     | `GO_MATCH_EXCLUDE`   | Exclude patterns (can be used multiple times) | -       |
| `-h, --help`    | -                    | Help for go-match                             | -       |
| `-v, --version` | -                    | Version for go-match                          | -       |

### Examples

```sh
# Filter Go files excluding a specific directory
go-match --include "**/*.go" --exclude "path/to/another/*" \
    path/to/dir1 \
    path/to/dir2 \
    path/to/file.go \
    path/to/another/file.go

# Output:
["path/to/file.go"]
```

### Pattern Syntax

- Use globstar patterns (`**`) to match multiple directory levels
- Use single asterisk (`*`) to match within a single directory level
- Patterns are case-sensitive
- Multiple include/exclude patterns can be specified

For detailed help:

```sh
go-match --help
```
