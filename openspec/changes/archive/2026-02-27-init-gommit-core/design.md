## Context

We are building a Go CLI tool. We need to interact with the local `git` binary and provide a simple user interface.

## Goals / Non-Goals

**Goals:**
- Reliable detection of staged files using `git status --porcelain`.
- Extensible structure for future Bubble Tea integration.

**Non-Goals:**
- Supporting multiple git repositories simultaneously (assume current directory).
- Complex commit message templates (fixed format for now).

## Decisions

### Decision 1: Use `os/exec` for Git commands

We will call the `git` binary directly rather than using a Go-native git library (like `go-git`). This is simpler for a prototype and more robust against local git configurations.

### Decision 2: Domain-driven Package Structure

We'll use an `internal/git` package for git operations. This keeps the core logic separate from the CLI (main.go).
