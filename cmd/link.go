package cmd

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/fatih/color"
    "github.com/spf13/cobra"
    "github.com/itzcodex24/git-guardian/internal/state"
)

var linkCmd = &cobra.Command{
    Use:   "link [folder]",
    Short: "Link a folder to Guardian for automatic backups",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        folder := args[0]
        absPath, err := filepath.Abs(folder)
        if err != nil {
            return fmt.Errorf("failed to resolve folder path: %w", err)
        }
        
        // Clean the path to remove any redundant separators
        absPath = filepath.Clean(absPath)

        info, err := os.Stat(absPath)
        if err != nil {
            return fmt.Errorf("folder does not exist: %w", err)
        }
        if !info.IsDir() {
            return fmt.Errorf("not a directory")
        }

        if err := state.AddFolder(absPath); err != nil {
            return fmt.Errorf("failed to link folder: %w", err)
        }

        color.Green("✓ Linked folder: %s", absPath)
        return nil
    },
}

func init() {
    rootCmd.AddCommand(linkCmd)
}
