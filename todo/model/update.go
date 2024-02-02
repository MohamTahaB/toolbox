package model

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/toolbox/todo/constants"
)

// Update function.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	cmd = nil
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If a width is set to the help menu, this will enable it to
		// truncate as needed.
		m.Help.Width = msg.Width

	// Now, on to the key messages.
	case tea.KeyMsg:
		switch {

		// Exit the program.
		case key.Matches(msg, constants.Keys.Quit):
			cmd = tea.Quit

		// Scroll the items, up.
		case key.Matches(msg, constants.Keys.Up):
			m.ListInfo.Selected--
			m.ListInfo.Selected += len(m.ListInfo.TasksList)
			m.ListInfo.Selected %= len(m.ListInfo.TasksList)

		// Scroll the items, down.
		case key.Matches(msg, constants.Keys.Down):
			m.ListInfo.Selected++
			m.ListInfo.Selected %= len(m.ListInfo.TasksList)

		// Select the highlighted item.
		case key.Matches(msg, constants.Keys.Check):
			m.ListInfo.TasksList[m.ListInfo.Selected].Done = !m.ListInfo.TasksList[m.ListInfo.Selected].Done

		// Toggle help
		case key.Matches(msg, constants.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll
		}

	}
	CommitListInfo(&m.ListInfo)
	return m, cmd
}
