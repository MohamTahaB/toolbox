package storage

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

// InitiateStorage checks if the storage directory is available.
// If not, it is initiated.
// Returns the dir to the JSON collection of present files, the dir to .md files, and an error.
func InitiateStorage() (string, string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", "", fmt.Errorf("error retrieving the current user: %v", err)
	}

	dir := filepath.Join(usr.HomeDir, ".toolbox", ".jotter", "mds")
	JSONDir := filepath.Join(usr.HomeDir, ".toolbox", ".jotter", "storage.json")

	// Check if both directories already exist.
	_, dirErr := os.Stat(dir)
	_, JSONErr := os.Stat(JSONDir)

	// Both JSON and .md dir exist.
	if !os.IsNotExist(dirErr) && !os.IsNotExist(JSONErr) {
		return JSONDir, dir, nil
	}

	// TODO! for now, if one of the dirs do not exist, everything will be set anew. This should change to accomodate all possible edge cases.

	// Create the dirs if one of them do not exist.
	if err = os.MkdirAll(filepath.Dir(dir), 0755); err != nil {
		return "", "", fmt.Errorf("error creating .md directory: %v", err)
	}

	if err = os.MkdirAll(filepath.Dir(JSONDir), 0755); err != nil {
		return "", "", fmt.Errorf("error creating JSON file: %v", err)
	}

	return JSONDir, dir, nil
}

// Create creates a new .md file.
func Create(name string) (string, error) {

	// Check if the storage is properly initiated.
	JSONDir, dir, err := InitiateStorage()
	if err != nil {
		return "", err
	}

	dir = filepath.Join(dir, fmt.Sprintf("%s.md", name))

	// Create the .md file.
	if _, err = os.Create(dir); err != nil {
		return "", fmt.Errorf("error creating the .md file: %v", err)
	}

	return dir, nil
}

// Pull pulls the content of a .md file name passed as a parameter.
func Pull(name string) (string, error) {

	// Check if the storage is properly initiated.
	dir, err := InitiateStorage()
	if err != nil {
		return "", err
	}

	dir = filepath.Join(dir, fmt.Sprintf("%s.md", name))

	// Open the file.
	b, err := os.ReadFile(dir)
	if err != nil {
		return "", fmt.Errorf("error reading the file: %v", err)
	}

	return string(b), nil
}

// Push overwrites the designated file with the content passed as parameter.
func Push(name, content string) error {

	// Check if the storage is properly initiated.
	dir, err := InitiateStorage()
	if err != nil {
		return err
	}

	dir = filepath.Join(dir, fmt.Sprintf("%s.md", name))

	if err = os.WriteFile(dir, []byte(content), 0755); err != nil {
		return fmt.Errorf("error pushing to the file: %v", err)
	}

	return nil
}
