package git

import (
    "bytes"
    "fmt"
    "os/exec"
    "strings"
    "time"
)

func hasChanges(folder string) (bool, error) {
    cmd := exec.Command("git", "-C", folder, "status", "--porcelain")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return false, fmt.Errorf("%w: %s", err, strings.TrimSpace(string(out)))
    }
    return len(bytes.TrimSpace(out)) > 0, nil
}

func CommitAndPush(folder string) error {
    changed, err := hasChanges(folder)
    if err != nil {
        return fmt.Errorf("git status failed: %w", err)
    }
    if !changed {
        return nil
    }

    if out, err := exec.Command("git", "-C", folder, "add", ".").CombinedOutput(); err != nil {
        return fmt.Errorf("git add failed: %v: %s", err, strings.TrimSpace(string(out)))
    }

    msg := "Auto backup " + time.Now().Format(time.RFC3339)
    if out, err := exec.Command("git", "-C", folder, "commit", "-m", msg).CombinedOutput(); err != nil {
        return fmt.Errorf("git commit failed: %v: %s", err, strings.TrimSpace(string(out)))
    }

    if out, err := exec.Command("git", "-C", folder, "push").CombinedOutput(); err != nil {
        return fmt.Errorf("git push failed: %v: %s", err, strings.TrimSpace(string(out)))
    }

    return nil
}
