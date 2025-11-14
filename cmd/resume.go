package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/itzcodex24/git-guardian/internal/supervisor"
)

var resumeCmd = &cobra.Command{
    Use:   "resume <id>",
    Args:  cobra.ExactArgs(1),
    Short: "Resume a paused watcher by id",
    RunE: func(cmd *cobra.Command, args []string) error {
        id := args[0]
        if err := supervisor.Resume(id); err != nil {
            return fmt.Errorf("resume failed: %w", err)
        }
        fmt.Println("Resumed:", id)
        return nil
    },
}

func init() {
    rootCmd.AddCommand(resumeCmd)
}
