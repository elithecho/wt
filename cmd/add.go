package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"worktrees/internal"
)

var addCmd = &cobra.Command{
	Use:     "add <path> [branch]",
	Aliases: []string{"create", "new"},
	Short:   "Create a new worktree",
	Long: `Create a new git worktree at the specified path.
If no branch is specified, a new branch will be created based on the folder name.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wm := internal.NewWorktreeManager()
		
		path := args[0]
		var branch string
		
		if len(args) > 1 {
			branch = args[1]
		} else {
			branch = filepath.Base(path)
		}
		
		fmt.Printf("Creating worktree at %s with branch %s...\n", path, branch)
		
		if err := wm.Add(path, branch); err != nil {
			fmt.Printf("Error creating worktree: %v\n", err)
			return
		}
		
		fmt.Printf("✅ Worktree created successfully!\n")
		fmt.Printf("   Path: %s\n", path)
		fmt.Printf("   Branch: %s\n", branch)
	},
}