package git

import (
	"fmt"
	"os/exec"
)

func StageAllChanges() error {
	cmd := exec.Command("git", "add", "--all")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to stage changes: %v\nOutput: %s", err, string(output))
	}

	return nil
}

func ExecuteCommit(commitMessage string) error {
	cmd := exec.Command("git", "commit", "-m", commitMessage)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to commit: %v\nOutput: %s", err, string(output))
	}

	return nil
}
