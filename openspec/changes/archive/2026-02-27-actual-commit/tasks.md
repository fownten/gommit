## 1. Git Execution Logic

- [x] 1.1 Implement `ExecuteCommit(message string) error` in `internal/git/status.go` (or a new file `internal/git/commit.go`).
- [x] 1.2 Update the `internal/git` package to handle output capture.

## 2. Integration

- [x] 2.1 Update `main.go` to call `git.ExecuteCommit` when the Bubble Tea model returns a confirmed state.
- [x] 2.2 Display the result (success or error) after the Bubble Tea program exits.

## 3. Verify

- [x] 3.1 Stage a change and run the tool to perform a real commit.
- [x] 3.2 Verify the commit appears in `git log`.
