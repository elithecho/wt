package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"worktrees/internal"
)

var switchCmd = &cobra.Command{
	Use:     "switch <name>",
	Aliases: []string{"cd", "goto", "go"},
	Short:   "Navigate to a worktree",
	Long: `Navigate to a worktree by name. This command will output the path
that can be used with shell integration to change directories.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wm := internal.NewWorktreeManager()
		name := args[0]
		
		wt, err := wm.FindByName(name)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			
			// Show available worktrees
			fmt.Println("\nAvailable worktrees:")
			worktrees, listErr := wm.List()
			if listErr == nil {
				for _, w := range worktrees {
					fmt.Printf("  - %s (%s)\n", filepath.Base(w.Path), w.Path)
				}
			}
			return
		}
		
		// Output the path for shell integration
		fmt.Print(wt.Path)
	},
}