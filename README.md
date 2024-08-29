# go-match Overview

`go-match` go-match is a simple command-line utility to check if a path matches a globstar pattern.

Returns a string "true" if the path matches the pattern, "false" otherwise.

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

    go-match [flags] <path> <pattern>

For more details on usage and configuration, run:

    go-match --help

This will display a comprehensive list of flags and their descriptions.
