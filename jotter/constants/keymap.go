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
	//TODO : check the update for the viewport.
}

func (km *KeyMap) WriteFileMode() {
	//TODO : To be precised later after settling on the solution of editiing the file : simple text area or form ?
}
