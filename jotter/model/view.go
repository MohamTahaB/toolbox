package model

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder
	var msg string

	// Render the appropriate section according to the state of the model.
	switch m.State {
	case ReadFileList:
		// Render the list of files.
		msg = m.List.View()
	case ReadFile:
		// Render the file viewport.
		msg = m.ViewPort.View()
	case WriteFile:
		// Render the form.
		msg = m.Form.View()
	}

	fmt.Fprintf(&b, "%s \n", msg)

	return ""
}
