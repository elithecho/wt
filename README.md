# worktrees (wt) 🌳

A friendly CLI tool for managing git worktrees with intuitive commands.

## Why?

Git worktrees are powerful but the commands aren't very intuitive. This tool provides a simple, user-friendly interface for common worktree operations.

## Installation

### Quick install (prebuilt binary):
```bash
curl -L https://github.com/elithecho/wt/releases/download/v0.1/worktree -o worktree
chmod +x worktree
sudo mv worktree /usr/local/bin/
```

### Or build from source:
```bash
git clone https://github.com/elithecho/wt.git
cd wt
make install
```

### Set up shell integration (required for directory changing):
```bash
worktree install
source ~/.zshrc  # or ~/.bashrc for bash
```

This installs the `wt` shell function that enables automatic directory changing.

## Commands

### `wt add <path> [branch]`
Create a new worktree at the specified path. If no branch is specified, creates a new branch based on the folder name.

```bash
wt add ../feature-auth        # Creates new branch "feature-auth"
wt add ../hotfix main         # Creates worktree from existing "main" branch
```

### `wt list` (aliases: `ls`, `l`)
List all worktrees with their paths, branches, and status.

```bash
wt list
```

### `wt switch <name>` (aliases: `s`, `cd`, `goto`, `go`)
Navigate to a worktree by name. Automatically changes directory with shell integration.

```bash
wt switch feature-auth
wt s feature-auth        # Short alias
```

### `wt remove <path>` (aliases: `rm`, `delete`, `del`)
Remove a worktree. The worktree must be clean (no uncommitted changes).

```bash
wt remove feature-auth
wt rm ../feature-auth
```

### `wt original` (aliases: `og`, `main`)
Navigate back to the main/original worktree.

```bash
wt og
```

### `wt clean` (aliases: `prune`)
Clean up stale worktree references for directories that no longer exist.

```bash
wt clean
```

## Shell Integration

After installation, run `worktree install` to automatically set up shell integration for your shell (bash/zsh/fish). This creates a `wt` function that intercepts navigation commands and handles directory changing automatically.

## Usage Examples

```bash
# Create a new feature worktree
wt add ../feature-login

# List all worktrees
wt list

# Switch to the feature worktree (automatically changes directory)
wt s feature-login

# Go back to main
wt og

# Remove the feature worktree when done
wt rm feature-login

# Clean up any stale references
wt clean
```

## Features

- 🚀 **Simple commands** - Intuitive interface over git worktree
- 📁 **Smart navigation** - Find worktrees by name, not full path
- 🔍 **Fuzzy matching** - Partial name matching for convenience
- 🏠 **Quick return** - Easy navigation back to main worktree
- 🧹 **Cleanup** - Automatic cleanup of stale references
- 🎨 **Pretty output** - Clean, readable status display

## License

MIT