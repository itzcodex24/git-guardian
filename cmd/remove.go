package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/itzcodex24/git-guardian/internal/supervisor"
)

var removeCmd = &cobra.Command{
    Use:   "remove <id>",
    Args:  cobra.ExactArgs(1),
    Short: "Remove a watcher configuration and stop it",
    RunE: func(cmd *cobra.Command, args []string) error {
        id := args[0]
        if err := supervisor.Remove(id); err != nil {
            return fmt.Errorf("remove failed: %w", err)
        }
        fmt.Println("Removed:", id)
        return nil
    },
}

func init() {
    rootCmd.AddCommand(removeCmd)
}
