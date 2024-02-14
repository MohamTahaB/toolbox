package model

import (
	"todo/constants"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

// Update function.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If a width is set to the help menu, this will enable it to
		// truncate as needed.
		m.Help.Width = msg.Width

	// Now, on to the key messages.
	case tea.KeyMsg:
		handleKeyMsg(&m, &msg, &cmd)
	}

	// Prep keymap according to the state of the model, and whether it can be navigated or not.
	switch m.State {
	case readingTasks:
		constants.Keys.ReadingMode(len(m.ListInfo.TasksList) == 0)
	case writingTasks:
		constants.Keys.WritingMode()
	case checkingDetails:
		constants.Keys.CheckingDetailsMode()
	case writingDetail:
		constants.Keys.WritingDetailMode()
	}

	CommitListInfo(&m.ListInfo)
	li, _ := RetrieveListInfo()
	m.ListInfo = *li
	return m, cmd
}

// Handles the updates to be done depending on the model's state.
func handleKeyMsg(m *Model, msg *tea.KeyMsg, cmd *tea.Cmd) {
	switch m.State {

	// Reading tasks state.
	case readingTasks:

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

			// In case the last element is deleted, the cursor should always stay in the scope of the new task list.
			if m.ListInfo.Selected >= len(m.ListInfo.TasksList) {
				m.ListInfo.Selected = len(m.ListInfo.TasksList) - 1
			}
			if m.ListInfo.Selected < 0 {
				m.ListInfo.Selected = 0
			}

			// Reprep the keymap in case the task list is empty.
			constants.Keys.ReadingMode(len(m.ListInfo.TasksList) == 0)

		// Toggle help.
		case key.Matches(*msg, constants.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll

		// Switch state.
		case key.Matches(*msg, constants.Keys.Write):
			m.State = writingTasks
			constants.Keys.WritingMode()

		// Get details description.
		case key.Matches(*msg, constants.Keys.Details):
			m.State = checkingDetails
			constants.Keys.CheckingDetailsMode()
		}

	// Writing tasks state.
	case writingTasks:

		// Focus on the text input field
		m.TaskInput.Focus()

		// Describe the behaviour for all key bindings.
		switch {
		case key.Matches(*msg, constants.Keys.Quit):
			m.TaskInput.Reset()
			// Toggle state.
			m.State = readingTasks
			constants.Keys.ReadingMode(len(m.ListInfo.TasksList) == 0)

		case key.Matches(*msg, constants.Keys.Check):
			// Create a new id and build the task struct.
			uuid, _ := uuid.NewV7()
			task := Task{
				ID:    uuid.String(),
				Title: m.TaskInput.Value(),
				Done:  false,
			}

			// Add the new task to the list info.
			m.ListInfo.TasksList = append(m.ListInfo.TasksList, task)

			// Toggle the state.
			m.State = readingTasks
			constants.Keys.ReadingMode(len(m.ListInfo.TasksList) == 0)

			// Reset the task input.
			m.TaskInput.Reset()

		default:
			m.TaskInput, *cmd = m.TaskInput.Update(*msg)
		}

	// Checking details state.
	case checkingDetails:

		// Describe the behaviour for all key bindings.
		switch {

		// Quit the details description.
		case key.Matches(*msg, constants.Keys.Quit):
			constants.Keys.ReadingMode(len(m.ListInfo.TasksList) == 0)
			m.State = readingTasks

		// Toggle help.
		case key.Matches(*msg, constants.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll

		// Edit description.
		case key.Matches(*msg, constants.Keys.Write):

			// Focus on the description input field.
			m.DescriptionInput.Focus()

			if len(m.ListInfo.TasksList[m.ListInfo.Selected].Description) != 0 {
				m.DescriptionInput.SetValue(m.ListInfo.TasksList[m.ListInfo.Selected].Description)
			}

			// Switch state to writing detail.
			m.State = writingDetail
			constants.Keys.WritingDetailMode()
			m.Help.ShowAll = true
		}

	// Writing description state.
	case writingDetail:

		// Describe the behaviour for all key bindings
		switch {
		case msg.Type == tea.KeyCtrlN:
			m.ListInfo.TasksList[m.ListInfo.Selected].Description = m.DescriptionInput.Value()

			// Reset description input.
			m.DescriptionInput.Reset()

			// Toggle state to checking details.
			m.State = checkingDetails

		case key.Matches(*msg, constants.Keys.Quit):
			// Reset description input.
			m.DescriptionInput.Reset()

			// Toggle state to checking details>
			m.State = checkingDetails

			constants.Keys.CheckingDetailsMode()

		default:
			m.DescriptionInput, *cmd = m.DescriptionInput.Update(*msg)
		}
	}
}
