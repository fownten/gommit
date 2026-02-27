## Context
The current implementation only calculates the message but doesn't persist it via git. We need a safe way to execute the command and relay output to the user.

## Goals / Non-Goals
**Goals:**
- Reliable execution of `git commit`.
- Clear feedback on success or failure.

**Non-Goals:**
- Complex interactive error resolution (e.g., editing hooks from within the tool).

## Decisions

### Decision 1: Use `os/exec` in a dedicated function
We will add `ExecuteCommit(message string) error` to the `internal/git` package. This maintains our package separation and makes the logic testable.

### Decision 2: Capture Stderr
We will capture the combined output (Stdout and Stderr) to ensure the user sees detailed error messages if a hook fails.

## Risks / Trade-offs
- **[Risk] Git Locks** → If another git process is running, the commit will fail. **Mitigation**: Standard error reporting will inform the user.
