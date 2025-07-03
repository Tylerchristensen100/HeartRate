package export

import (
	"os"
	"strings"
)

func createDirectoryIfNotExists(path string) error {
	// Check if the path has a trailing filename
	if len(path) > 0 && path[len(path)-1] != '/' {
		// If it does, strip the filename to get the directory path
		endOfPath := strings.LastIndex(path, "/")
		path = path[:endOfPath]
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}
