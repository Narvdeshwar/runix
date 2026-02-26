package setup

import (
	"os"
	"path/filepath"
)

// EnsureDirs creates the necessary local directories for myterm in the user's home directory.
func EnsureDirs() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	base := filepath.Join(home, ".myterm")

	dirs := []string{"data", "state", "sockets", "workspace"}
	for _, d := range dirs {
		path := filepath.Join(base, d)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
