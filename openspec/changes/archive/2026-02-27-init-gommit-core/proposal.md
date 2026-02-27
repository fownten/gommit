## Why

We want a standardized way to create commit messages using a Go-based CLI. This change establishes the project foundation and the core logic for detecting staged changes.

## What Changes

- Initialize Go module `github.com/fownten/gommit`.
- Implement a `git` package to fetch staged files and their status (New/Modified/Deleted).
- Create a basic CLI entry point that prompts for type (feat/fix/chore) and description.

## Capabilities

### New Capabilities
- `detect-staged-files`: Programmatically list files staged for commit.
- `generate-commit-message`: Compose a formatted string based on user input.

### Modified Capabilities

## Impact

- `go.mod`: Project initialization.
- `internal/git/status.go`: Logic for interacting with git.
- `main.go`: CLI entry point.
