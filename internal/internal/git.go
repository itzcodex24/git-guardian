package git

import (
    "bytes"
    "fmt"
    "os/exec"
    "strings"
    "time"
)

// hasChanges returns true if there are staged or unstaged changes.
func hasChanges(folder string) (bool, error) {
    cmd := exec.Command("git", "-C", folder, "status", "--porcelain")
    out, err := cmd.Output()
    if err != nil {
        return false, err
    }
    return len(bytes.TrimSpace(out)) > 0, nil
}

// CommitAndPush adds, commits (if changes), and pushes.
// It will skip commits if there are no changes.
func CommitAndPush(folder string) error {
    changed, err := hasChanges(folder)
    if err != nil {
        return fmt.Errorf("git status failed: %w", err)
    }
    if !changed {
        // nothing to commit
        return nil
    }

    // git add .
    if out, err := exec.Command("git", "-C", folder, "add", ".").CombinedOutput(); err != nil {
        return fmt.Errorf("git add failed: %v: %s", err, strings.TrimSpace(string(out)))
    }

    // commit
    msg := "Auto backup " + time.Now().Format(time.RFC3339)
    if out, err := exec.Command("git", "-C", folder, "commit", "-m", msg).CombinedOutput(); err != nil {
        // ignore "nothing to commit" or other benign messages
        return fmt.Errorf("git commit failed: %v: %s", err, strings.TrimSpace(string(out)))
    }

    // push
    if out, err := exec.Command("git", "-C", folder, "push").CombinedOutput(); err != nil {
        return fmt.Errorf("git push failed: %v: %s", err, strings.TrimSpace(string(out)))
    }

    fmt.Println("[guardian] backup pushed at", time.Now().Format(time.RFC3339))
    return nil
}
