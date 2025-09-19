package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wt",
	Short: "A friendly CLI tool for managing git worktrees",
	Long: `worktrees (wt) is a CLI tool that makes working with git worktrees intuitive.
It provides simple commands to create, list, navigate, and manage git worktrees.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(switchCmd)
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(originalCmd)
	rootCmd.AddCommand(installCmd)
}