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
        fmt.Printf("\n%-8s %-15s %-38s %-10s %-25s\n", "ID", "Mode", "Folder", "Status", "LastRun")
        fmt.Println(strings.Repeat("-", 110))
        for _, w := range list {
            status := "paused"
            if !w.Paused {
                status = "running"
            }
            mode := w.Mode
            if w.Mode == "interval" && w.Interval != "" {
                mode = fmt.Sprintf("interval (%s)", w.Interval)
            } else if w.Mode == "watch" && w.Debounce != "" {
                mode = fmt.Sprintf("watch (%s)", w.Debounce)
            }
            fmt.Printf("%-8s %-15s %-38s %-10s %-25s\n", w.ID, mode, w.Folder, status, w.LastRun)
        }
        return nil
    },
}

func init() {
    rootCmd.AddCommand(listenersCmd)
}
