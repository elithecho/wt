package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"worktrees/internal"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "List all worktrees",
	Long:    `Display all git worktrees with their paths, branches, and status.`,
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
		
		fmt.Println("Worktrees:")
		for i, wt := range worktrees {
			icon := "📁"
			if i == 0 {
				icon = "🏠" // Main worktree
			}
			
			status := ""
			if wt.Bare {
				status += " (bare)"
			}
			if wt.Locked {
				status += " (locked)"
			}
			
			branch := wt.Branch
			if branch == "" {
				branch = "detached HEAD"
			}
			
			fmt.Printf("  %s %s\n", icon, filepath.Base(wt.Path))
			fmt.Printf("     Path: %s\n", wt.Path)
			fmt.Printf("     Branch: %s%s\n", branch, status)
			if i < len(worktrees)-1 {
				fmt.Println()
			}
		}
	},
}