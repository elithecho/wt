package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"worktrees/internal"
)

var originalCmd = &cobra.Command{
	Use:     "original",
	Aliases: []string{"og", "main"},
	Short:   "Navigate to the main worktree",
	Long:    `Navigate back to the original/main worktree directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		wm := internal.NewWorktreeManager()
		
		worktrees, err := wm.List()
		if err != nil {
			fmt.Printf("Error listing worktrees: %v\n", err)
			return
		}
		
		if len(worktrees) == 0 {
			fmt.Println("No worktrees found")
			return
		}
		
		// The first worktree is always the main/original one
		mainWorktree := worktrees[0]
		
		// Output the path for shell integration
		fmt.Print(mainWorktree.Path)
	},
}