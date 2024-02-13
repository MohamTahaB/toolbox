package model

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

type state int

const (
	readingTasks state = iota
	writingTasks
	checkingDetails
	writingDetail
)

// Model structure, consists of a list of all tasks, and the index of the selected one.
type Model struct {
	State            state
	TaskInput        textinput.Model
	DescriptionInput textarea.Model
	ListInfo         ListInfo
	Help             help.Model
}


type ListInfo struct {
	TasksList []Task `json:"tasksList"`
	Selected  int    `json:"selected"`
}

// Task structure, consists of a title and a boolean that describes its state.
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func NewModel() (*Model, error) {
	// Initiate description text area
	ta := textarea.New()
	ta.Placeholder = "Wagwan"

	// Retrieve list info from local storage.
	li, err := RetrieveListInfo()
	if err != nil {
		return nil, err
	}
	return &Model{
		State:            readingTasks,
		ListInfo:         *li,
		Help:             help.New(),
		TaskInput:        textinput.New(),
		DescriptionInput: ta,
	}, nil
}
