package cmd

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/yourorg/myterm/internal/runtime"
	"github.com/yourorg/myterm/internal/tui"
)

// dashboardCmd represents the dashboard command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Launch the immersive myterm TUI dashboard",
	Long:  `Opens a high-end terminal interface for managing workspaces and AI agents in an independent window.`,
	Run: func(cmd *cobra.Command, args []string) {
		isChild := false
		for _, v := range os.Args {
			if v == "--child" {
				isChild = true
				break
			}
		}

		if !isChild {
			exe, _ := os.Executable()
			// Use powershell to start a new window safely
			psArgs := fmt.Sprintf("Start-Process -FilePath '%s' -ArgumentList 'dashboard', '--child'", exe)
			newCmd := exec.Command("powershell.exe", "-Command", psArgs)
			if err := newCmd.Start(); err != nil {
				fmt.Printf("Error opening new window: %v\n", err)
			}
			return
		}

		p := tea.NewProgram(tui.InitialModel(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running TUI: %v\n", err)
			os.Exit(1)
		}

		// After quitting TUI, if we want to launch a shell:
		if err := runtime.CreateWorkspace(); err != nil {
			fmt.Printf("Failed to launch workspace: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dashboardCmd)
}
