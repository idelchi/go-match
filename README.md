# go-match

`go-match` is a simple command-line utility that filters a list of paths based on include and exclude globstar patterns.

The program takes a list of paths as input and outputs a JSON array of paths that match
the include patterns and don't match the exclude patterns.

## Getting Started

### Prerequisites

- Go 1.23 or higher
- Git

### Installation

Clone the repository and build the binary with:

    git clone ssh://git@code.swisscom.com:2222/swisscom/scsa-shared-tools/go-match.git
    cd go-match
    go build -o go-match .

Alternatively, you can install it directly using:

    go install code.swisscom.com/swisscom/scsa-shared-tools/go-match@latest

### Usage

    go-match [flags] [paths...]

The available flags include:

    --include: Specify one or more include patterns (can be used multiple times)
    --exclude: Specify one or more exclude patterns (can be used multiple times)

Example:

    go-match --include "**/*.go" --exclude "path/to/another/*" path/to/dir1 path/to/dir2 path/to/file.go path/to/another/file.go

will output

    ["path/to/file.go"]

For more details on usage and configuration, run:

    go-match --help

This will display a comprehensive list of flags and their descriptions.

All flags can be set through environment variables. The prefix _GO_MATCH_ is used to avoid conflicts.
For example, to set the include paths, use `GO_MATCH_INCLUDE`.
