//go:build windows

package runtime

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// StartContainerd on Windows is a no-op because we use WSL directly.
func StartContainerd() (*exec.Cmd, error) {
	// We check if WSL is available
	_, err := exec.LookPath("wsl.exe")
	if err != nil {
		return nil, fmt.Errorf("WSL2 is not installed or not in PATH: %v", err)
	}
	return nil, nil
}

// ImportImage on Windows uses wsl --import.
func ImportImage() error {
	// First, check if the distro already exists
	// We use 'wsl --status' or just 'wsl --list' and check for the name.
	// To avoid UTF-16 issues, we can try to filter via findstr in cmd
	checkCmd := exec.Command("cmd.exe", "/c", "wsl.exe --list | findstr myterm-workspace")
	if err := checkCmd.Run(); err == nil {
		return nil // Found it, already exists
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	base := filepath.Join(home, ".myterm")
	distroDir := filepath.Join(base, "distro")
	imagePath := filepath.Join("images", "workspace.tar")

	// Create distro directory if it doesn't exist
	os.MkdirAll(distroDir, 0755)

	// Command: wsl --import myterm-workspace <InstallLocation> <FileName>
	cmd := exec.Command("wsl.exe", "--import", "myterm-workspace", distroDir, imagePath, "--version", "2")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// CreateWorkspace on Windows launches the WSL distro in a NEW window.
func CreateWorkspace() error {
	// We use 'cmd.exe /c start' to spawn a new process window
	// 'wsl.exe -d myterm-workspace' is the command to run in that window
	cmd := exec.Command("cmd.exe", "/c", "start", "wsl.exe", "-d", "myterm-workspace")

	// Since we are starting a detached window, we don't bridge Stdin/Stdout here
	return cmd.Run()
}
