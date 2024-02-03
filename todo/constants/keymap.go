package constants

import "github.com/charmbracelet/bubbles/key"

// Help KeyMap struct.
type KeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Help   key.Binding
	Quit   key.Binding
	Check  key.Binding
	Write  key.Binding
	Delete key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Help, k.Quit,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Check, k.Delete, k.Write},
		{k.Help, k.Quit},
	}
}
