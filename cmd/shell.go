package cmd

import (
	"fmt"
	"time"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/yourorg/myterm/internal/runtime"
	"github.com/yourorg/myterm/internal/setup"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Open an isolated Linux workspace",
	Long:  `Launches a contained Linux environment where you can install AI tools safely.`,
	Run: func(cmd *cobra.Command, args []string) {
		pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("MYTERM - AI Workspace")

		pterm.Info.Println("Initializing your isolated Linux environment...")

		// 1. Ensure directories exist
		spinner, _ := pterm.DefaultSpinner.Start("Checking host environment...")
		if err := setup.EnsureDirs(); err != nil {
			spinner.Fail(fmt.Sprintf("Failed to initialize directories: %v", err))
			return
		}
		spinner.Success("Host environment ready!")

		// 2. Start containerd/Verify WSL
		spinner, _ = pterm.DefaultSpinner.Start("Starting container runtime...")
		_, err := runtime.StartContainerd()
		if err != nil {
			spinner.Fail(fmt.Sprintf("Failed to start runtime: %v", err))
			return
		}
		spinner.Success("Runtime active!")

		// 3. Import workspace image
		spinner, _ = pterm.DefaultSpinner.Start("Checking workspace image...")
		if err := runtime.ImportImage(); err != nil {
			spinner.Warning(fmt.Sprintf("Image check note: %v", err))
		} else {
			spinner.Success("Workspace image imported!")
		}

		// 4. Create and enter workspace
		pterm.Success.Println("Launcher complete. Dropping into shell...")
		time.Sleep(1 * time.Second)

		if err := runtime.CreateWorkspace(); err != nil {
			pterm.Error.Printf("Failed to launch workspace: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}
