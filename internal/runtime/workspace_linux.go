//go:build linux

package runtime

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// ImportImage loads the pre-built workspace image into containerd.
func ImportImage() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	socket := filepath.Join(home, ".myterm", "sockets", "containerd.sock")
	imagePath := filepath.Join("images", "workspace.tar")

	cmd := exec.Command("sudo", "./runtime/ctr",
		"--address", socket,
		"images", "import", imagePath,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// CreateWorkspace runs the workspace container with an interactive TTY.
func CreateWorkspace() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	base := filepath.Join(home, ".myterm")
	socket := filepath.Join(base, "sockets", "containerd.sock")
	workspaceDir := filepath.Join(base, "workspace")

	// We use ctr to run the container. In a more advanced version, we'd use the containerd Go client.
	cmd := exec.Command("sudo", "./runtime/ctr",
		"--address", socket,
		"run",
		"--rm",
		"-t",
		"--mount", fmt.Sprintf("type=bind,src=%s,dst=/workspace,options=rbind:rw", workspaceDir),
		"docker.io/library/myterm/workspace:latest",
		"workspace-shell",
		"bash",
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
