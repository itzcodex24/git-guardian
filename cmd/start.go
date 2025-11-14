package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/itzcodex24/git-guardian/internal/state"
)

var (
    startInterval string
    startDebounce string
    startWatch    bool
)

var startCmd = &cobra.Command{
    Use:   "start <folder>",
    Args:  cobra.ExactArgs(1),
    Short: "Start a watcher for a linked folder (watch mode or interval mode)",
    RunE: func(cmd *cobra.Command, args []string) error {
        folder := args[0]
        list := state.Get()
        for i := range list {
            if list[i].Folder == folder {
                if startWatch {
                    list[i].Mode = "watch"
                    list[i].Debounce = startDebounce
                } else if startInterval != "" {
                    list[i].Mode = "interval"
                    list[i].Interval = startInterval
                } else {
                    return fmt.Errorf("must specify --watch or --interval <duration>")
                }
                list[i].Paused = false
                state.Update(list)
                fmt.Println("Watcher activated for:", folder)
                return nil
            }
        }
        return fmt.Errorf("folder not linked: %s. Use `guardian link` first", folder)
    },
}

func init() {
    startCmd.Flags().BoolVar(&startWatch, "watch", false, "Use file-change watch mode")
    startCmd.Flags().StringVar(&startInterval, "interval", "", "Use interval mode, e.g. 5m")
    startCmd.Flags().StringVar(&startDebounce, "debounce", "30s", "Debounce duration for watch mode")
    rootCmd.AddCommand(startCmd)
}
