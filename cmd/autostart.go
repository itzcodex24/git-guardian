package cmd

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"

    "github.com/spf13/cobra"
    "github.com/itzcodex24/git-guardian/internal/supervisor"
)

var autostartCmd = &cobra.Command{
    Use:   "autostart",
    Short: "Manage launch-at-login (launchd) autostart for guardian",
}

var autostartEnableCmd = &cobra.Command{
    Use:   "enable",
    Short: "Enable autostart (install launch agent and load it)",
    RunE: func(cmd *cobra.Command, args []string) error {
        plist := supervisor.GenerateLaunchdPlist()
        home, _ := os.UserHomeDir()
        plistDir := filepath.Join(home, "Library", "LaunchAgents")
        if err := os.MkdirAll(plistDir, 0755); err != nil {
            return fmt.Errorf("failed to make launchagents dir: %w", err)
        }
        plistPath := filepath.Join(plistDir, "com.gitguardian.agent.plist")
        if err := os.WriteFile(plistPath, []byte(plist), 0644); err != nil {
            return fmt.Errorf("failed to write plist: %w", err)
        }
        exec.Command("launchctl", "load", plistPath).Run()
        fmt.Println("Autostart enabled. Plist:", plistPath)
        return nil
    },
}

var autostartDisableCmd = &cobra.Command{
    Use:   "disable",
    Short: "Disable autostart (unload and remove launch agent)",
    RunE: func(cmd *cobra.Command, args []string) error {
        home, _ := os.UserHomeDir()
        plistPath := filepath.Join(home, "Library", "LaunchAgents", "com.gitguardian.agent.plist")
        exec.Command("launchctl", "unload", plistPath).Run()
        os.Remove(plistPath)
        fmt.Println("Autostart disabled. Removed:", plistPath)
        return nil
    },
}

func init() {
    autostartCmd.AddCommand(autostartEnableCmd)
    autostartCmd.AddCommand(autostartDisableCmd)
    rootCmd.AddCommand(autostartCmd)
}
