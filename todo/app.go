package todo

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/toolbox/todo/model"
)

func App() {
	m, err := model.NewModel()
	if err != nil {
		fmt.Printf("there has been an error: %v", err)
		os.Exit(1)
	}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("there has been an error: %v", err)
		os.Exit(1)
	}
}
