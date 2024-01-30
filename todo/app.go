package todo

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	helpmenu "github.com/toolbox/todo/helpMenu"
	"github.com/toolbox/todo/model"
)

func App() {
	m, err := model.NewModel(Keys)
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

var Keys = helpmenu.KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move down"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Check: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↪", "toggle check / uncheck"),
	),
}
