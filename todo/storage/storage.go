package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/toolbox/todo/model"
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

	// Create the dir if it does not exist.
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

// RetrieveModel serves as a call to get the content of the JSON file, that each transaction will do, similarly to a SQL migrate.
func RetrieveModel() (*model.Model, error) {
	fileName, err := InitiateStorage()

	// Check if there was an error initiating the storage.
	if err != nil {
		return nil, err
	}

	// Read the file.
	b, err := os.ReadFile(fileName)
	// Check if reading the file went error free.
	if err != nil {
		return nil, fmt.Errorf("error reading the file: %v", err)
	}

	// Initiate a ptr to the model.
	var m *model.Model
	// Unmarshal the byte slice.
	if err = json.Unmarshal(b, m); err != nil {
		return nil, fmt.Errorf("error unmarshalling the file: %v", err)
	}

	return m, nil
}
