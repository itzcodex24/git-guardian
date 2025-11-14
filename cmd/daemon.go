package cmd

import (
    "github.com/spf13/cobra"
    "github.com/itzcodex24/git-guardian/internal/supervisor"
)

var daemonCmd = &cobra.Command{
    Use:    "daemon",
    Hidden: true,
    Short:  "Internal: start daemon (used by launchd)",
    RunE: func(cmd *cobra.Command, args []string) error {
        supervisor.StartAllBlocking()
        return nil
    },
}

func init() {
    rootCmd.AddCommand(daemonCmd)
}
