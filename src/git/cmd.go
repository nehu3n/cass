package git

import (
	"fmt"
	"os/exec"
)

func HasChanges() (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")

	output, err := cmd.CombinedOutput()

	if err != nil {
		return false, fmt.Errorf("failed to check git status: %v\nOutput: %s", err, string(output))
	}

	return len(output) != 0, nil
}

func HasPendingChanges() (bool, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")

	output, err := cmd.CombinedOutput()

	if err != nil {
		return false, fmt.Errorf("failed to check git status: %v\nOutput: %s", err, string(output))
	}

	return len(output) != 0, nil
}

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
