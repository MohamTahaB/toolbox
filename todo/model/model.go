package model

import (
	"github.com/charmbracelet/bubbles/help"
)

// Model structure, consists of a list of all tasks, and the index of the selected one.
type Model struct {
	ListInfo ListInfo
	Help     help.Model
}

type ListInfo struct {
	TasksList []Task `json:"tasksList"`
	Selected  int    `json:"selected"`
}

// Task structure, consists of a title and a boolean that describes its state.
type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func NewModel() (*Model, error) {
	li, err := RetrieveListInfo()
	if err != nil {
		return nil, err
	}
	return &Model{
		ListInfo: *li,
		Help:     help.New(),
	}, nil
}
