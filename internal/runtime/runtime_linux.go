//go:build linux

package runtime

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// StartContainerd launches the bundled containerd daemon using sudo.
func StartContainerd() (*exec.Cmd, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	base := filepath.Join(home, ".myterm")
	socket := filepath.Join(base, "sockets", "containerd.sock")

	// In a real production scenario, we would check if the binary exists in ./runtime/
	// For this MVP, we assume Phase 2 assets are placed correctly later.
	cmd := exec.Command("sudo", "./runtime/containerd",
		"--address", socket,
		"--root", filepath.Join(base, "data"),
		"--state", filepath.Join(base, "state"),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	// Wait for containerd to boot and create the socket
	time.Sleep(2 * time.Second)
	return cmd, nil
}
