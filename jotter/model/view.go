package model

import (
	"fmt"
	"strings"
	"toolbox/jotter/constants"
)

func (m Model) View() string {
	var b strings.Builder
	var msg string

	// Render the appropriate section according to the state of the model.
	switch m.State {
	case ReadFileList:
		// Render the list of files.
		if len(m.List.Items()) == 0 {
			msg = "No files available for the moment"
		} else {
			msg = m.List.View()
		}
	case ReadFile:
		// Render the file viewport.
		msg = m.ViewPort.View()
	case WriteFile:
		// Render the form.
		msg = m.Form.View()
	}

	fmt.Fprintf(&b, "%s \n", msg)

	fmt.Fprintf(&b, "%s", m.Help.View(constants.HelpKeyMap))

	return b.String()
}
