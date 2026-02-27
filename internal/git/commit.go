package git

import (
	"fmt"
	"os/exec"
)

// ExecuteCommit executes the actual git commit command with the given message.
// It returns an error if the command fails, including the output from git.
func ExecuteCommit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if len(output) > 0 {
			return fmt.Errorf("git commit failed: %w\nOutput: %s", err, string(output))
		}
		return fmt.Errorf("git commit failed: %w", err)
	}
	return nil
}
