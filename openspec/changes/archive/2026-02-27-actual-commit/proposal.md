## Why
Currently, the application only simulates the commit process. To be a functional tool, it must execute the actual `git commit` command using the generated message when the user confirms.

## What Changes
- Replace the simulated commit log with a real `os/exec` call to `git commit -m "<message>"`.
- Handle potential git errors (e.g., hooks failing, lock files) and display them to the user.

## Capabilities

### New Capabilities
- `execute-git-commit`: Logic to invoke the git binary with the final message.

### Modified Capabilities

## Impact
- `main.go`: Update the confirmation logic to call the execution function.
- `internal/git/`: Add an `ExecuteCommit` function.
