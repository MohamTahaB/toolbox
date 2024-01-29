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
func InitiateStorage() (*StorageFile, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("error retrieving the current user: %v", err)
	}

	fileName := filepath.Join(usr.HomeDir, ".toolbox", ".todo", "storage.json")

	// Check if the JSON file already exists.
	if _, err := os.Stat(fileName); os.IsExist(err) {
		return &StorageFile{
			Path: fileName,
		}, nil
	}

	// Create the dir if it does not exist.
	if err = os.MkdirAll(filepath.Dir(fileName), 0755); err != nil {
		return nil, fmt.Errorf("error creating directory: %v", err)
	}

	file, createErr := os.Create(fileName)
	// Checks if there was an error when creating the JSON file.
	if createErr != nil {
		return nil, fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()
	return &StorageFile{
		Path: fileName,
	}, nil
}

// RetrieveModel serves as a call to get the content of the JSON file, that each transaction will do, similarly to a SQL migrate.
func RetrieveModel() (*model.Model, error) {

	// First, initiate the storage.
	storageFile, err := InitiateStorage()

	// Check if there was an error initiating the storage.
	if err != nil {
		return nil, err
	}

	// Read the file.
	b, err := os.ReadFile(storageFile.GetPath())
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

// CommitModel commits a model instance to the storage JSON file.
func CommitModel(m *model.Model) error {

	// First, initiate the storage.
	storageFile, err := InitiateStorage()
	if err != nil {
		return err
	}

	// Marshal the model.
	b, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("error marshalling the file: %v", err)
	}

	// Write marshalled model in the storage JSON file.
	if err = os.WriteFile(storageFile.GetPath(), b, 0755); err != nil {
		return fmt.Errorf("error writing to the file: %v", err)
	}

	return nil
}
