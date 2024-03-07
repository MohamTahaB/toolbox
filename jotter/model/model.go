package model

import (
	"toolbox/jotter/storage"
	fileinfo "toolbox/jotter/storage/fileInfo"
	filelist "toolbox/jotter/storage/fileList"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/huh"
)

type state int

const (
	ReadFileList state = iota
	ReadFile
	WriteFile
)

type Model struct {
	State       state
	List        list.Model
	Form        huh.Form
	FileID      string
	CurrentFile fileinfo.FileInfo
	ViewPort    viewport.Model
	Help        help.Model
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
	for id, file := range *fileMap {
		filesSlice = append(filesSlice, filelist.FileItem{
			ID: id,
			FileTitle: file.FileTitle,
			FileDesc: file.FileDesc,
		})
	}

	// Initiate the list and the file edit.
	m.List = list.New(filesSlice, list.NewDefaultDelegate(), 0, 0)

	// Get the first id of the files in the file map.
	for id := range *fileMap {
		m.FileID = id
		break
	}

	// Initiate help.
	m.Help = help.New()

	m.List.Title = "Jotter"
	m.CurrentFile.InitiateFile()
	m.InitiateForm()

	return &m, nil

}

// Initiates the form in the model, and links its inputs to the model's current file.
func (m *Model) InitiateForm() {
	m.Form = *huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is the title gonna be ?").
				Prompt(">>>").
				Value(&m.CurrentFile.Title),

			huh.NewText().
				Title("Let your fingertips go wild !").
				Value(&m.CurrentFile.Content),
		),
	)
}
