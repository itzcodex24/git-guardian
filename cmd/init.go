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
    Short: "Create a new GitHub repository using the GitHub CLI (`gh`) and initialize locally",
    RunE: func(cmd *cobra.Command, args []string) error {
        rd := bufio.NewReader(os.Stdin)

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

        fmt.Println("Created GitHub repo:", name)
        return nil
    },
}

func init() {
    rootCmd.AddCommand(initCmd)
}
