package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Worktree struct {
	Path   string `json:"worktree"`
	Head   string `json:"HEAD"`
	Branch string `json:"branch"`
	Bare   bool   `json:"bare"`
	Locked bool   `json:"locked"`
}

type WorktreeManager struct{}

func NewWorktreeManager() *WorktreeManager {
	return &WorktreeManager{}
}

func (wm *WorktreeManager) List() ([]Worktree, error) {
	cmd := exec.Command("git", "worktree", "list", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list worktrees: %w", err)
	}

	return parseWorktreeList(string(output))
}

func (wm *WorktreeManager) Add(path, branch string) error {
	args := []string{"worktree", "add"}
	
	if branch != "" {
		args = append(args, "-b", branch)
	}
	
	args = append(args, path)
	
	if branch != "" && !strings.HasPrefix(branch, "-b") {
		args = append(args, branch)
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func (wm *WorktreeManager) Remove(path string) error {
	cmd := exec.Command("git", "worktree", "remove", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func (wm *WorktreeManager) Prune() error {
	cmd := exec.Command("git", "worktree", "prune", "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func (wm *WorktreeManager) FindByName(name string) (*Worktree, error) {
	worktrees, err := wm.List()
	if err != nil {
		return nil, err
	}

	for _, wt := range worktrees {
		if filepath.Base(wt.Path) == name {
			return &wt, nil
		}
		if strings.Contains(wt.Path, name) {
			return &wt, nil
		}
	}

	return nil, fmt.Errorf("worktree '%s' not found", name)
}

func parseWorktreeList(output string) ([]Worktree, error) {
	var worktrees []Worktree
	var current Worktree
	
	lines := strings.Split(strings.TrimSpace(output), "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		if line == "" {
			if current.Path != "" {
				worktrees = append(worktrees, current)
				current = Worktree{}
			}
			continue
		}
		
		parts := strings.SplitN(line, " ", 2)
		key := parts[0]
		
		switch key {
		case "worktree":
			if len(parts) > 1 {
				current.Path = parts[1]
			}
		case "HEAD":
			if len(parts) > 1 {
				current.Head = parts[1]
			}
		case "branch":
			if len(parts) > 1 {
				current.Branch = strings.TrimPrefix(parts[1], "refs/heads/")
			}
		case "bare":
			current.Bare = true
		case "locked":
			current.Locked = true
		}
	}
	
	if current.Path != "" {
		worktrees = append(worktrees, current)
	}
	
	return worktrees, nil
}