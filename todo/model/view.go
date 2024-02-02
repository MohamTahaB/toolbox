package model

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/toolbox/todo/constants"
)

// View function.
func (m Model) View() string {

	var b strings.Builder
	var check, cursor string

	// Build the string.
	fmt.Fprint(&b, "Your ToDo list...\n\n")

	if len(m.ListInfo.TasksList) == 0 {
		fmt.Fprint(&b, "no tasks found, please add some...")
	} else {

		for index, task := range m.ListInfo.TasksList {

			// Configure task style.
			taskStyle := lipgloss.NewStyle()

			// Configure the cursor.
			if index == m.ListInfo.Selected {
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
	}

	// Build the help view.
	helpView := m.Help.View(constants.Keys)
	fmt.Fprintf(&b, "\n\n%s\n\n", helpView)

	// Check if the state is in reading or writing.
	if m.State == writing {
		fmt.Fprintf(&b, "$ %s", m.TaskInput.View())
	}

	// Render the final string.
	return b.String()

}
