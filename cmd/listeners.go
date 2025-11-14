package cmd

import (
    "fmt"
    "strings"

    "github.com/fatih/color"
    "github.com/spf13/cobra"
    "github.com/itzcodex24/git-guardian/internal/supervisor"
)

var listenersCmd = &cobra.Command{
    Use:   "listeners",
    Short: "List all configured watchers",
    RunE: func(cmd *cobra.Command, args []string) error {
        list := supervisor.List()
        
        cyan := color.New(color.FgCyan, color.Bold)
        fmt.Println()
        cyan.Printf("%-8s %-15s %-38s %-10s %-25s\n", "ID", "Mode", "Folder", "Status", "LastRun")
        color.Cyan(strings.Repeat("─", 110))
        
        for _, w := range list {
            statusColor := color.New(color.FgYellow)
            statusText := "paused"
            if !w.Paused {
                statusColor = color.New(color.FgGreen)
                statusText = "running"
            }
            
            mode := w.Mode
            modeColor := color.New(color.FgWhite)
            if w.Mode == "interval" && w.Interval != "" {
                mode = fmt.Sprintf("interval (%s)", w.Interval)
                modeColor = color.New(color.FgMagenta)
            } else if w.Mode == "watch" && w.Debounce != "" {
                mode = fmt.Sprintf("watch (%s)", w.Debounce)
                modeColor = color.New(color.FgBlue)
            }
            
            idColor := color.New(color.FgHiWhite, color.Bold)
            idColor.Printf("%-8s ", w.ID)
            modeColor.Printf("%-15s ", mode)
            fmt.Printf("%-38s ", w.Folder)
            statusColor.Printf("%-10s ", statusText)
            color.New(color.FgHiBlack).Printf("%-25s\n", w.LastRun)
        }
        return nil
    },
}

func init() {
    rootCmd.AddCommand(listenersCmd)
}
