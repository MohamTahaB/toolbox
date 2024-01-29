package todo

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/toolbox/todo/model"
)

func App() {
	m, err := model.RetrieveModel()
	if err != nil {
		fmt.Printf("there has been an error: %v", err)
		os.Exit(1)
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("there has been an error: %v", err)
		os.Exit(1)
	}
}
