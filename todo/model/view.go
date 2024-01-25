package model

import (
	"fmt"
	"strings"
)

// View function.
func (m Model) View() string {
	var b strings.Builder
	var check, cursor string

	// Build the string.
	fmt.Fprint(&b, "Your ToDo list...\n\n")
	for index, task := range m.TasksList {

		// Configure the cursor.
		if index == m.Selected {
			cursor = "   >>>"
		} else {
			cursor = "      "
		}

		// Configure the check.
		if task.Done {
			check = SelectedStyle.Render("✓")
		} else {
			check = " "
		}

		// Add the task to the builder.
		fmt.Fprintf(&b, "%s [%s] %s;\n", cursor, check, task.Title)
	}
	return b.String()

}
