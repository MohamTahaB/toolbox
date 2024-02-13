package constants

import (
	"github.com/charmbracelet/bubbles/key"
)

// Help KeyMap struct.
type KeyMap struct {
	Up      key.Binding
	Down    key.Binding
	Help    key.Binding
	Quit    key.Binding
	Check   key.Binding
	Write   key.Binding
	Delete  key.Binding
	Details key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Check, k.Help, k.Quit,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Details},
		{k.Check, k.Delete, k.Write},
		{k.Help, k.Quit},
	}
}

// Disables the navigation buttons when the tasks list is empty,
// avoiding any nav errors.
func (k *KeyMap) DisableNav() {
	navKeys := k.NavKeys()
	for _, keyPtr := range navKeys {
		(*keyPtr).SetEnabled(false)
	}
}

// Enables the navigation buttons.
// To be called upon whenever the tasks list is not empty.
func (k *KeyMap) EnableNav() {
	for _, key := range k.NavKeys() {
		key.SetEnabled(true)
	}
}

// Returns a nav keys ptrs slice.
func (k *KeyMap) NavKeys() []*key.Binding {
	return []*key.Binding{&k.Up, &k.Down, &k.Delete, &k.Check}
}

// Preps the keymap when the state is writing.
func (k *KeyMap) WritingMode() {

	k.DisableNav()
	k.Write.SetEnabled(false)
	k.Check.SetEnabled(true)
	k.Check.SetHelp("↪", "validate the task")
}

// Preps the keymap when the state is reading.
func (k *KeyMap) ReadingMode(isTaskListEmpty bool) {

	k.Write.SetEnabled(true)
	k.Check.SetEnabled(true)
	k.Check.SetHelp("↪", "toggle check / uncheck")
	k.Write.SetHelp("n | w", "add new task")

	if isTaskListEmpty {
		k.DisableNav()
	} else {
		k.EnableNav()
	}
}

// Preps the keymap when the state is checking details of a certain task in a list.
func (k *KeyMap) CheckingDetailsMode() {
	k.DisableNav()
	k.Write.SetEnabled(true)
	k.Write.SetHelp("n | w", "edit description")
}

// Preps the keymap when the state is writing detail of a certain task in a list.
func (k *KeyMap) WritingDetailMode() {
	k.DisableNav()
	k.Write.SetEnabled(false)

}