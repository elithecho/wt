package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"worktrees/internal"
)

var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"prune"},
	Short:   "Clean up stale worktree references",
	Long:    `Remove stale administrative files for worktrees that no longer exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		wm := internal.NewWorktreeManager()
		
		fmt.Println("Cleaning up stale worktree references...")
		
		if err := wm.Prune(); err != nil {
			fmt.Printf("Error cleaning worktrees: %v\n", err)
			return
		}
		
		fmt.Println("✅ Cleanup completed!")
	},
}