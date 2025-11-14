package cmd

import (
    "fmt"
    "strings"

    "github.com/spf13/cobra"
    "github.com/itzcodex24/git-guardian/internal/supervisor"
)

var listenersCmd = &cobra.Command{
    Use:   "listeners",
    Short: "List all configured watchers",
    RunE: func(cmd *cobra.Command, args []string) error {
        list := supervisor.List()
        fmt.Printf("\n%-8s %-10s %-45s %-10s %-25s\n", "ID", "Mode", "Folder", "Status", "LastRun")
        fmt.Println(strings.Repeat("-", 110))
        for _, w := range list {
            status := "paused"
            if !w.Paused {
                status = "running"
            }
            fmt.Printf("%-8s %-10s %-45s %-10s %-25s\n", w.ID, w.Mode, w.Folder, status, w.LastRun)
        }
        return nil
    },
}

func init() {
    rootCmd.AddCommand(listenersCmd)
}
