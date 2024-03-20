package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	filelist "toolbox/jotter/storage/fileList"
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
	// The mds dir is enough as it is the most ramified one.
	if err = os.MkdirAll(filepath.Clean(dir), 0755); err != nil {
		return "", "", fmt.Errorf("error creating .md directory: %v", err)
	}

	// Create the JSON file.
	if _, JSONErr = os.Create(JSONDir); JSONErr != nil {
		return "", "", fmt.Errorf("error creating JSON file: %v", JSONErr)
	}

	// Populate the json file for the first time.
	var IDmap map[string]filelist.FileDesignation
	b, err := json.Marshal(&IDmap)
	if err != nil {
		return "", "", fmt.Errorf("error initiating the JSON file: %v", err)
	}

	if err = os.WriteFile(JSONDir, b, 0755); err != nil {
		return "", "", fmt.Errorf("error initiating the JSON file: %v", err)
	}

	return JSONDir, dir, nil
}

// Create creates a new .md file.
func Create(id string) (string, error) {

	// Check if the storage is properly initiated.
	_, dir, err := InitiateStorage()
	if err != nil {
		return "", err
	}

	dir = filepath.Join(dir, fmt.Sprintf("%s.md", id))

	// Create the .md file.
	if _, err = os.Create(dir); err != nil {
		return "", fmt.Errorf("error creating the .md file: %v", err)
	}

	return dir, nil
}

// Pull pulls the content of a .md file name passed as a parameter.
func Pull(id string) (string, error) {

	// Check if the storage is properly initiated.
	_, dir, err := InitiateStorage()
	if err != nil {
		return "", err
	}

	dir = filepath.Join(dir, fmt.Sprintf("%s.md", id))

	// Open the file.
	b, err := os.ReadFile(dir)
	if err != nil {
		return "", fmt.Errorf("error reading the file: %v", err)
	}

	return string(b), nil
}

// Push overwrites the designated file with the content passed as parameter.
func Push(id, content string) error {

	// Check if the storage is properly initiated.
	_, dir, err := InitiateStorage()
	if err != nil {
		return err
	}

	dir = filepath.Join(dir, fmt.Sprintf("%s.md", id))

	if err = os.WriteFile(dir, []byte(content), 0755); err != nil {
		return fmt.Errorf("error pushing to the file: %v", err)
	}

	return nil
}
