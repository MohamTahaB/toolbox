package help

import "github.com/charmbracelet/bubbles/key"

// Help KeyMap struct.
type KeyMap struct {
	Up key.Binding
	Down key.Binding
	Help key.Binding
	Quit key.Binding
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Help, k.Quit},
	}
}