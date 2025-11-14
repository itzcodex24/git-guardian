package cmd

import (
    "fmt"
    "os"

    "github.com/google/uuid"
    "github.com/spf13/cobra"
    "github.com/itzcodex24/git-guardian/internal/state"
)

var linkCmd = &cobra.Command{
    Use:   "link <folder>",
    Args:  cobra.ExactArgs(1),
    Short: "Link an existing folder to guardian (adds it to persistent state, initially paused)",
    RunE: func(cmd *cobra.Command, args []string) error {
        folder := args[0]
        if _, err := os.Stat(folder); os.IsNotExist(err) {
            return fmt.Errorf("folder does not exist: %s", folder)
        }

        w := state.WatcherState{
            ID:     uuid.NewString()[:8],
            Folder: folder,
            Mode:   "none",
            Paused: true,
        }
        state.Append(w)
        fmt.Println("Linked folder (paused):", folder, "id:", w.ID)
        fmt.Println("Run `guardian start <folder> --watch` or `--interval <duration>` to enable.")
        return nil
    },
}

func init() {
    rootCmd.AddCommand(linkCmd)
}
