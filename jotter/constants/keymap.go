package constants

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Quit   key.Binding
	Enter  key.Binding
	Create key.Binding
	Help   key.Binding
}

var HelpKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↪", "Enter"),
	),
	Create: key.NewBinding(
		key.WithKeys("n", "w"),
		key.WithHelp("n | w", "Create new ..."),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "Toggle help"),
	),
}

// Resets the keymap to its default case, meaning that all key bindings will be enabled.
// Should be called before any call of the state mode. Thus avoiding any issues with suddently dissappearing help key bindings when switching states.
func (km *KeyMap) ResetToDefault() {
	//TODO! check if it can be done in a better and tidyier way.
	(*km).Up.SetEnabled(true)
	(*km).Down.SetEnabled(true)
	(*km).Quit.SetEnabled(true)
	(*km).Enter.SetEnabled(true)
	(*km).Create.SetEnabled(true)
	(*km).Help.SetEnabled(true)
}

// Disables all the keymap bindings.
func (km *KeyMap) DisableAll() {
	(*km).Up.SetEnabled(false)
	(*km).Down.SetEnabled(false)
	(*km).Quit.SetEnabled(false)
	(*km).Enter.SetEnabled(false)
	(*km).Create.SetEnabled(false)
	(*km).Help.SetEnabled(false)
}

func (km *KeyMap) ReadFileListMode() {

	// Reset to default before acting on the key bindings.
	km.ResetToDefault()

	// No need for the bindings, except for quitting, creating, help and entering, as the rest will be handled by the native update func of the list component.
	(*km).Up.SetEnabled(false)
	(*km).Down.SetEnabled(false)
}

func (km *KeyMap) WriteFileListMode() {
	//TODO : check the update for the form.
}

func (km *KeyMap) ReadFileMode() {

	// Needs quit and help only, the rest is in the viewport update.
	km.DisableAll()
	(*km).Help.SetEnabled(true)
	(*km).Quit.SetEnabled(true)
}

func (km *KeyMap) WriteFileMode() {
	//In this case, we are going through huh, therefore, all the keymap should be disabled.
	km.DisableAll()
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Enter, k.Quit, k.Help,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Enter},
		{k.Create, k.Help, k.Quit},
	}
}
