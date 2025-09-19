package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"worktrees/internal"
)

var removeCmd = &cobra.Command{
	Use:     "remove <path>",
	Aliases: []string{"rm", "delete", "del"},
	Short:   "Remove a worktree",
	Long:    `Remove a git worktree. The worktree must be clean (no uncommitted changes).`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wm := internal.NewWorktreeManager()
		path := args[0]
		
		// Try to find by name first
		if wt, err := wm.FindByName(path); err == nil {
			path = wt.Path
		}
		
		fmt.Printf("Removing worktree: %s\n", path)
		
		if err := wm.Remove(path); err != nil {
			fmt.Printf("Error removing worktree: %v\n", err)
			return
		}
		
		fmt.Println("✅ Worktree removed successfully!")
	},
}