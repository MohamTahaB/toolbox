package model

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/toolbox/todo/constants"
)

// Update function.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	cmd = nil

	if len(m.ListInfo.TasksList) == 0 {
		constants.Keys.DisableNav()
	} else {
		constants.Keys.EnableNav()
	}

	if m.State == writing {
		constants.Keys.Check.SetEnabled(true)
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If a width is set to the help menu, this will enable it to
		// truncate as needed.
		m.Help.Width = msg.Width

	// Now, on to the key messages.
	case tea.KeyMsg:
		switch m.State {
		case reading:

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

			// Delete the highlighted item.
			case key.Matches(msg, constants.Keys.Delete):
				if err := m.ListInfo.TasksList[m.ListInfo.Selected].DeleteTask(); err != nil {
					panic(err)
				}
				li, err := RetrieveListInfo()
				if err != nil {
					panic(err)
				}
				m.ListInfo = *li

			// Toggle help.
			case key.Matches(msg, constants.Keys.Help):
				m.Help.ShowAll = !m.Help.ShowAll

			// Switch state.
			case key.Matches(msg, constants.Keys.Write):
				m.State = writing
			}
		case writing:

			m.TaskInput.Focus()
			switch {
			case key.Matches(msg, constants.Keys.Quit):
				m.TaskInput = textinput.New()
				m.State = reading

			case key.Matches(msg, constants.Keys.Check):
				uuid, _ := uuid.NewV7()
				task := Task{
					ID:    uuid.String(),
					Title: m.TaskInput.View(),
					Done:  false,
				}

				m.ListInfo.TasksList = append(m.ListInfo.TasksList, task)

				m.State = reading

			default:
				m.TaskInput, cmd = m.TaskInput.Update(msg)
			}

		}

	}
	CommitListInfo(&m.ListInfo)
	li, _ := RetrieveListInfo()
	m.ListInfo = *li
	return m, cmd
}
