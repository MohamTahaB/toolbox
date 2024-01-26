package storage

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

// InitiateStorage checks for the JSON file, creates it if necessary.
// Returns the JSON directory to be manipulated later by transactors, and an error if there are issues.
func InitiateStorage() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("error retrieving the current user: %v", err)
	}

	fileName := filepath.Join(usr.HomeDir, ".toolbox", ".todo", "storage.json")

	// Create the dir if does not exist.
	if err = os.MkdirAll(filepath.Dir(fileName), 0755); err != nil {
		return "", fmt.Errorf("error creating directory: %v", err)
	}

	// Create or open the hidden JSON file.
	file, err := os.Create(fileName)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()
	return fileName, nil
}
