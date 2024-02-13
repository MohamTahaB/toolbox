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

	switch {
	case m.State == checkingDetails:
		m.checkDetails(&b)
	case m.State == writingDetail:
		m.writeDetail(&b)
	// Default case.
	default:
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

	}

	// Build the help view.
	helpView := m.Help.View(constants.Keys)
	fmt.Fprintf(&b, "\n\n%s\n\n", helpView)

	// Check if the state is in reading or writing.
	if m.State == writingTasks {
		fmt.Fprintf(&b, "$ %s", m.TaskInput.View())
	}

	// Render the final string.
	return b.String()

}

// checkDetails renders the details of current task.
// It is called upon when the status is checkingDetails.
func (m Model) checkDetails(b *strings.Builder) {
	// Add the title.
	fmt.Fprintf(b, "Title: %s\n", m.ListInfo.TasksList[m.ListInfo.Selected].Title)

	// Add the description of the highlighted task.
	desc := m.ListInfo.TasksList[m.ListInfo.Selected].Description

	if len(desc) == 0 {
		desc = "no description available."
	}
	fmt.Fprintf(b, "Description: %s\n", desc)

	// Add status.
	status := "Done"
	if !m.ListInfo.TasksList[m.ListInfo.Selected].Done {
		status = "Pending"
	}

	fmt.Fprintf(b, "Status: %s\n", status)
}

// writeDetail renders the info of the current task, as well as the description input.
func (m Model) writeDetail(b *strings.Builder) {
	m.checkDetails(b)
	
	// Add the description text area view.
	fmt.Fprintf(b, "\n\n%s\n", m.DescriptionInput.View())
}
