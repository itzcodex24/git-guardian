package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/itzcodex24/git-guardian/internal/supervisor"
)

var pauseCmd = &cobra.Command{
    Use:   "pause <id>",
    Args:  cobra.ExactArgs(1),
    Short: "Pause a running watcher by id",
    RunE: func(cmd *cobra.Command, args []string) error {
        id := args[0]
        if err := supervisor.Pause(id); err != nil {
            return fmt.Errorf("pause failed: %w", err)
        }
        fmt.Println("Paused:", id)
        return nil
    },
}

func init() {
    rootCmd.AddCommand(pauseCmd)
}
