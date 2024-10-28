# go-match

`go-match` is a simple command-line utility that filters a list of paths based on include and exclude globstar patterns.

The program takes a list of paths as input and outputs a JSON array of paths that match
the include patterns and don't match the exclude patterns.

## Installation

### From source

```sh
go install github.com/idelchi/go-match/cmd/go-match@latest
```

### From installation script

```sh
curl -sSL https://raw.githubusercontent.com/idelchi/go-match/refs/heads/main/install.sh | sh -s -- -d ~/.local/bin
```

## Usage

```sh
go-match [flags] [paths...]
```

Run `go-match` with the desired flags. The available flags include:

```sh
--include: Specify one or more include patterns (can be used multiple times)
--exclude: Specify one or more exclude patterns (can be used multiple times)
```

Example:

```sh
go-match --include "**/*.go" --exclude "path/to/another/*" path/to/dir1 path/to/dir2 path/to/file.go path/to/another/file.go
```

will output

```sh
["path/to/file.go"]
```

For more details on usage and configuration, run:

```sh
go-match --help
```

This will display a comprehensive list of flags and their descriptions.
