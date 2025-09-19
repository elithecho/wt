package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install wt shell function for automatic directory changing",
	Long: `Install a 'wt' shell function that intercepts navigation commands
and automatically changes directories. Other commands are passed through 
to the worktree binary.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := installShellIntegration(); err != nil {
			fmt.Printf("Error installing shell integration: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Shell integration installed successfully!")
		fmt.Println("Please restart your shell or run 'source <config_file>' to use the wt function.")
		fmt.Println("\nUsage:")
		fmt.Println("  wt og            - Jump to original worktree")
		fmt.Println("  wt s <name>      - Jump to worktree by name")
		fmt.Println("  wt switch <name> - Jump to worktree by name")
		fmt.Println("  wt <other>       - Pass command to worktree binary")
	},
}

func detectShell() (string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "", fmt.Errorf("unable to detect shell from $SHELL environment variable")
	}
	
	shellName := filepath.Base(shell)
	return shellName, nil
}

func getShellConfigPath(shell string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to get home directory: %v", err)
	}

	switch shell {
	case "bash":
		// Try .bashrc first, fall back to .bash_profile
		bashrc := filepath.Join(homeDir, ".bashrc")
		if _, err := os.Stat(bashrc); err == nil {
			return bashrc, nil
		}
		return filepath.Join(homeDir, ".bash_profile"), nil
	case "zsh":
		return filepath.Join(homeDir, ".zshrc"), nil
	case "fish":
		configDir := filepath.Join(homeDir, ".config", "fish")
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return "", fmt.Errorf("unable to create fish config directory: %v", err)
		}
		return filepath.Join(configDir, "config.fish"), nil
	default:
		return "", fmt.Errorf("unsupported shell: %s", shell)
	}
}

func getBashZshIntegration() string {
	return `
# wt shell function - intercepts navigation commands for directory changing
wt() {
    case "$1" in
        "og"|"original"|"main")
            # Navigate to original worktree
            local path=$(worktree og 2>/dev/null)
            if [ $? -eq 0 ] && [ -n "$path" ]; then
                builtin cd "$path" 2>/dev/null || cd "$path"
            else
                worktree og
            fi
            ;;
        "s"|"switch"|"cd"|"goto"|"go")
            # Navigate to specified worktree
            if [ $# -lt 2 ]; then
                worktree switch
            else
                local path=$(worktree switch "$2" 2>/dev/null)
                if [ $? -eq 0 ] && [ -n "$path" ]; then
                    builtin cd "$path" 2>/dev/null || cd "$path"
                else
                    worktree switch "$2"
                fi
            fi
            ;;
        *)
            # Pass everything else to worktree binary
            worktree "$@"
            ;;
    esac
}
`
}

func getFishIntegration() string {
	return `
# wt shell function - intercepts navigation commands for directory changing
function wt
    switch $argv[1]
        case "og" "original" "main"
            # Navigate to original worktree
            set path (worktree og 2>/dev/null)
            if test $status -eq 0 -a -n "$path"
                builtin cd $path 2>/dev/null; or cd $path
            else
                worktree og
            end
        case "s" "switch" "cd" "goto" "go"
            # Navigate to specified worktree
            if test (count $argv) -lt 2
                worktree switch
            else
                set path (worktree switch $argv[2] 2>/dev/null)
                if test $status -eq 0 -a -n "$path"
                    builtin cd $path 2>/dev/null; or cd $path
                else
                    worktree switch $argv[2]
                end
            end
        case "*"
            # Pass everything else to worktree binary
            worktree $argv
    end
end
`
}

func installShellIntegration() error {
	shell, err := detectShell()
	if err != nil {
		return err
	}

	fmt.Printf("Detected shell: %s\n", shell)

	configPath, err := getShellConfigPath(shell)
	if err != nil {
		return err
	}

	fmt.Printf("Shell config file: %s\n", configPath)

	// Read existing config
	var existingContent []byte
	if _, err := os.Stat(configPath); err == nil {
		existingContent, err = os.ReadFile(configPath)
		if err != nil {
			return fmt.Errorf("unable to read config file: %v", err)
		}
	}

	// Check if integration is already installed
	existingContentStr := string(existingContent)
	if strings.Contains(existingContentStr, "wt shell function - intercepts navigation commands") {
		fmt.Println("Shell integration is already installed.")
		return nil
	}

	// Get the appropriate integration code
	var integration string
	switch shell {
	case "bash", "zsh":
		integration = getBashZshIntegration()
	case "fish":
		integration = getFishIntegration()
	default:
		return fmt.Errorf("unsupported shell: %s", shell)
	}

	// Append integration to config file
	file, err := os.OpenFile(configPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("unable to open config file for writing: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(integration); err != nil {
		return fmt.Errorf("unable to write integration to config file: %v", err)
	}

	return nil
}