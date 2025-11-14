package cmd

import (
    "fmt"
    "os"

    "github.com/itzcodex24/git-guardian/internal/supervisor"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "guardian",
    Short: "Git Guardian — automatic git-backed folder watcher for macOS",
    Long:  "Guardian watches folders, auto-commits and pushes changes to git/GitHub and runs as a macOS background agent.",
}

func Execute() {
    supervisor.StartAll()

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
