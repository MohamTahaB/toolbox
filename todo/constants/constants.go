package constants

import (
	"github.com/charmbracelet/bubbles/key"
	helpmenu "github.com/toolbox/todo/helpMenu"
)

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
	Delete: key.NewBinding(
		key.WithKeys("delete"),
		key.WithHelp("⌦", "delete task"),
	),
}
