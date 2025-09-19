# worktrees (wt) 🌳

A friendly CLI tool for managing git worktrees with intuitive commands.

## Why?

Git worktrees are powerful but the commands aren't very intuitive. This tool provides a simple, user-friendly interface for common worktree operations.

## Installation

```bash
go install github.com/yourusername/worktrees@latest
```

Or build from source:
```bash
git clone https://github.com/yourusername/worktrees.git
cd worktrees
go build -o wt main.go
sudo mv wt /usr/local/bin/
```

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

### `wt switch <name>` (aliases: `cd`, `goto`, `go`)
Navigate to a worktree by name. Use with shell integration for automatic directory changing.

```bash
wt switch feature-auth
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

To enable automatic directory changing, add this to your shell configuration:

### Bash/Zsh
```bash
wt_cd() {
    local path=$(wt switch "$1" 2>/dev/null)
    if [ $? -eq 0 ] && [ -n "$path" ]; then
        cd "$path"
    else
        wt switch "$1"
    fi
}

wt_og() {
    local path=$(wt og 2>/dev/null)
    if [ $? -eq 0 ] && [ -n "$path" ]; then
        cd "$path"
    else
        wt og
    fi
}

alias wtcd=wt_cd
alias wtog=wt_og
```

### Fish
```fish
function wt_cd
    set path (wt switch $argv[1] 2>/dev/null)
    if test $status -eq 0 -a -n "$path"
        cd $path
    else
        wt switch $argv[1]
    end
end

function wt_og
    set path (wt og 2>/dev/null)
    if test $status -eq 0 -a -n "$path"
        cd $path
    else
        wt og
    end
end

alias wtcd=wt_cd
alias wtog=wt_og
```

## Usage Examples

```bash
# Create a new feature worktree
wt add ../feature-login

# List all worktrees
wt list

# Switch to the feature worktree
wtcd feature-login

# Go back to main
wtog

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