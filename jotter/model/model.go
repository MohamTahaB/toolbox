package model

import (
	"toolbox/jotter/storage"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
)

type state int

const (
	ReadFileList state = iota
	WriteFileList
	ReadFile
	WriteFile
)

type Model struct {
	State    state
	List     list.Model
	FileEdit textarea.Model
	ViewPort viewport.Model
}

// Initiates the app model.
func InitiateModel() (*Model, error) {
	var m Model

	// Initiate the storage.
	JSONDir, _, err := storage.InitiateStorage()
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON file.
	fileMap, err := storage.Unmarshal(JSONDir)
	if err != nil {
		return nil, err
	}

	// Build list items slice.
	var filesSlice []list.Item
	for _, file := range *fileMap {
		filesSlice = append(filesSlice, file)
	}

	// Initiate the list and the file edit.
	m.List = list.New(filesSlice, list.NewDefaultDelegate(), 0, 0)

	m.List.Title = "Jotter"

	return &m, nil

}
