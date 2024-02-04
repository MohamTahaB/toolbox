package model

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/toolbox/todo/constants"
)

// Update function.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil

	// Prep keymap according to the state of the model, and whether it can be navigated or not.
	switch m.State {
	case reading:
		constants.Keys.ReadingMode(len(m.ListInfo.TasksList) == 0)
	case writing:
		// Show all help so that the interrogation point can be used in the task title.
		m.Help.ShowAll = true
		constants.Keys.WritingMode()
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If a width is set to the help menu, this will enable it to
		// truncate as needed.
		m.Help.Width = msg.Width

	// Now, on to the key messages.
	case tea.KeyMsg:
		handleKeyMsg(&m, &msg, &cmd)
	}
	CommitListInfo(&m.ListInfo)
	li, _ := RetrieveListInfo()
	m.ListInfo = *li
	return m, cmd
}

// Handles the updates to be done depending on whether the model is in reading or writing states.
func handleKeyMsg(m *Model, msg *tea.KeyMsg, cmd *tea.Cmd) {
	switch m.State {

	// Reading state.
	case reading:

		// Describe the behaviour for all key bindings.
		switch {

		// Exit the program.
		case key.Matches(*msg, constants.Keys.Quit):
			*cmd = tea.Quit

		// Scroll the items, up.
		case key.Matches(*msg, constants.Keys.Up):
			m.ListInfo.Selected--
			m.ListInfo.Selected += len(m.ListInfo.TasksList)
			m.ListInfo.Selected %= len(m.ListInfo.TasksList)

		// Scroll the items, down.
		case key.Matches(*msg, constants.Keys.Down):
			m.ListInfo.Selected++
			m.ListInfo.Selected %= len(m.ListInfo.TasksList)

		// Select the highlighted item.
		case key.Matches(*msg, constants.Keys.Check):
			m.ListInfo.TasksList[m.ListInfo.Selected].Done = !m.ListInfo.TasksList[m.ListInfo.Selected].Done

		// Delete the highlighted item.
		case key.Matches(*msg, constants.Keys.Delete):
			if err := m.ListInfo.TasksList[m.ListInfo.Selected].DeleteTask(); err != nil {
				panic(err)
			}
			li, err := RetrieveListInfo()
			if err != nil {
				panic(err)
			}
			m.ListInfo = *li

		// Toggle help.
		case key.Matches(*msg, constants.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll

		// Switch state.
		case key.Matches(*msg, constants.Keys.Write):
			m.State = writing
			constants.Keys.WritingMode()
		}

	// Writing state.
	case writing:

		// Focus on the text input field
		m.TaskInput.Focus()

		// Describe the behaviour for all key bindings.
		switch {
		case key.Matches(*msg, constants.Keys.Quit):
			m.TaskInput.Reset()
			// Toggle state.
			m.State = reading
			constants.Keys.ReadingMode(len(m.ListInfo.TasksList) == 0)

		case key.Matches(*msg, constants.Keys.Check):
			// Create a new id and build the task struct.
			uuid, _ := uuid.NewV7()
			task := Task{
				ID:    uuid.String(),
				Title: m.TaskInput.View(),
				Done:  false,
			}

			// Add the new task to the list info.
			m.ListInfo.TasksList = append(m.ListInfo.TasksList, task)

			// Toggle the state.
			m.State = reading
			constants.Keys.ReadingMode(len(m.ListInfo.TasksList) == 0)

		default:
			m.TaskInput, *cmd = m.TaskInput.Update(msg)
		}

	}
}
