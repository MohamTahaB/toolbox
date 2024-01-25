package model

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// View function.
func (m Model) View() string {
	var b strings.Builder
	var check, cursor string

	// Build the string.
	fmt.Fprint(&b, "Your ToDo list...\n\n")
	for index, task := range m.TasksList {

		// Configure task style.
		taskStyle := lipgloss.NewStyle()

		// Configure the cursor.
		if index == m.Selected {
			cursor = CursorStyle.Render("   >>>")
			taskStyle = taskStyle.Inherit(SelectedStyle)
		} else {
			cursor = "      "
		}

		// Configure the check.
		if task.Done {
			check = SelectedStyle.Render("✓")
			taskStyle = taskStyle.Inherit(StrikeThroughStyle)
		} else {
			check = " "
		}

		// Add the task to the builder.
		fmt.Fprintf(&b, "%s [%s] %s;\n", cursor, check, taskStyle.Render(task.Title))
	}
	return b.String()

}
