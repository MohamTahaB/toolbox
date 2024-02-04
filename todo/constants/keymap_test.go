package constants_test

import (
	"testing"

	"github.com/charmbracelet/bubbles/key"
	"github.com/toolbox/todo/constants"
)

// Checks if disable nav does, indeed, disable nav.
func TestDisableNav_OK(t *testing.T) {

	// Create a key map with some nav key bindings.
	keyMap := constants.KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("up", "navigate up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("down", "navigate down"),
		),
		Delete: key.NewBinding(
			key.WithKeys("delete"),
			key.WithHelp("delete", "delete task"),
		),
		Check: key.NewBinding(
			key.WithKeys("c"),
			key.WithHelp("c", "check task"),
		),
	}

	keyMap.DisableNav()

	if keyMap.Up.Enabled() || keyMap.Down.Enabled() || keyMap.Delete.Enabled() || keyMap.Check.Enabled() {
		t.Error("unexpected outcome: nav keys should have been disabled")
	}

}

// Checks whether enable nav does, indeed, enable nav.
func TestEnableNav_OK(t *testing.T) {

	// Create a key map with some nav key bindings.
	keyMap := constants.KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("up", "navigate up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("down", "navigate down"),
		),
		Delete: key.NewBinding(
			key.WithKeys("delete"),
			key.WithHelp("delete", "delete task"),
		),
		Check: key.NewBinding(
			key.WithKeys("c"),
			key.WithHelp("c", "check task"),
		),
	}

	// Disable all the nav key bindings.
	for _, key := range keyMap.NavKeys() {
		key.SetEnabled(false)
	}

	// Enable nav keys.
	keyMap.EnableNav()

	if !(keyMap.Up.Enabled() && keyMap.Down.Enabled() && keyMap.Delete.Enabled() && keyMap.Check.Enabled()) {
		t.Error("unexpected outcome: nav keys should have been enabled")
	}

}
