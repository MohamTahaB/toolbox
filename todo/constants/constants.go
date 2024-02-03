package constants

import (
	"github.com/charmbracelet/bubbles/key"
)

var Keys = KeyMap{
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
	Write: key.NewBinding(
		key.WithKeys("n", "w"),
		key.WithHelp("n | w", "add new task"),
	),
	Delete: key.NewBinding(
		key.WithKeys("backspace"),
		key.WithHelp("⌦", "delete task"),
	),
}
