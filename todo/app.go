package todo

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/toolbox/todo/model"
)

func App() {
	m := model.Model{
		TasksList: []model.Task{
			{
				Title: "Make the laundry",
				Done:  false,
			},
			{
				Title: "Wash the dishes",
				Done:  false,
			},
			{
				Title: "Tidy the room",
				Done:  false,
			},
		},
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("there has been an error: %v", err)
		os.Exit(1)
	}
}
