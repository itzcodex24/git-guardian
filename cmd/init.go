package cmd

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"

    "github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
    Use:   "init",
    Short: "Initialize Guardian in this folder (uses existing Git repo if present)",
    RunE: func(cmd *cobra.Command, args []string) error {
        rd := bufio.NewReader(os.Stdin)

        if _, err := os.Stat(".git"); err == nil {
            fmt.Println("This folder already has a Git repository.")
            fmt.Print("Guardian will use the existing repository. Continue? [Y/n]: ")
            resp, _ := rd.ReadString('\n')
            resp = strings.ToLower(strings.TrimSpace(resp))
            if resp == "n" {
                fmt.Println("Aborting init.")
                return nil
            }
            fmt.Println("Guardian will use the existing Git repository for automatic backups.")
            return nil
        }

        fmt.Print("Repository name (user/repo or repo): ")
        name, _ := rd.ReadString('\n')
        name = strings.TrimSpace(name)

        fmt.Print("Private? (y/N): ")
        privRaw, _ := rd.ReadString('\n')
        priv := strings.ToLower(strings.TrimSpace(privRaw)) == "y"

        flag := "--public"
        if priv {
            flag = "--private"
        }

        create := exec.Command("gh", "repo", "create", name, flag, "--confirm")
        create.Stdout = os.Stdout
        create.Stderr = os.Stderr
        if err := create.Run(); err != nil {
            return fmt.Errorf("failed to create repo via gh: %w", err)
        }

        addRemote := exec.Command("git", "remote", "add", "origin", fmt.Sprintf("git@github.com:%s.git", name))
        addRemote.Run()

        status := exec.Command("git", "status", "--porcelain")
        out, _ := status.Output()
        if len(out) > 0 {
            exec.Command("git", "add", ".").Run()
            exec.Command("git", "commit", "-m", "Initial commit by Guardian").Run()
            exec.Command("git", "push", "-u", "origin", "main").Run()
        }

        fmt.Println("Guardian setup complete. GitHub repository linked and ready for backups.")
        return nil
    },
}

func init() {
    rootCmd.AddCommand(initCmd)
}
