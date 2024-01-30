package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Update function.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		// Exit the program.
		case "ctrl+c", "esc", "q":
			cmd = tea.Quit

		// Scroll the items, up.
		case "up":
			m.ListInfo.Selected--
			m.ListInfo.Selected += len(m.ListInfo.TasksList)
			m.ListInfo.Selected %= len(m.ListInfo.TasksList)

		// Scroll the items, down.
		case "down":
			m.ListInfo.Selected++
			m.ListInfo.Selected %= len(m.ListInfo.TasksList)

		// Select the highlighted item.
		case "enter":
			m.ListInfo.TasksList[m.ListInfo.Selected].Done = !m.ListInfo.TasksList[m.ListInfo.Selected].Done
		}

	}
	CommitListInfo(&m.ListInfo)
	return m, cmd
}
