package model

import tea "github.com/charmbracelet/bubbletea"

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
			m.Selected--
			m.Selected += len(m.TasksList)
			m.Selected %= len(m.TasksList)

		// Scroll the items, down.
		case "down":
			m.Selected++
			m.Selected %= len(m.TasksList)

		// Select the highlighted item.
		case "enter":
			m.TasksList[m.Selected].Done = !m.TasksList[m.Selected].Done
		}

	}
	return m, cmd
}
