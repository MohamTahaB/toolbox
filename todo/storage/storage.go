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

	// Check if the JSON file already exists.
	if _, err := os.Stat(fileName); os.IsExist(err) {
		return fileName, nil
	}

	// Create the dir when it does not exist.
	if err = os.MkdirAll(filepath.Dir(fileName), 0755); err != nil {
		return "", fmt.Errorf("error creating directory: %v", err)
	}

	file, createErr := os.Create(fileName)
	// Checks if there was an error when creating the JSON file.
	if createErr != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()
	return fileName, nil
}
