package model

import (
	"toolbox/jotter/constants"
	filelist "toolbox/jotter/storage/fileList"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m Model) Update(msg tea.Msg) (tea.Msg, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	// The sole change is the size of the window.
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil

	// Case when the msg is a key msg.
	case tea.KeyMsg:
		handleKeyMsg(&m, &msg, &cmd)
	}

	// Handle the keymap mode.

	return nil, nil
}

func handleKeyMsg(m *Model, msg *tea.KeyMsg, cmd *tea.Cmd) {
	switch m.State {

	// Case when the state is reading the file list.
	case ReadFileList:
		switch {
		// Exit the program.
		case key.Matches(*msg, constants.HelpKeyMap.Quit):
			*cmd = tea.Quit

		// Enter to a list item.
		case key.Matches(*msg, constants.HelpKeyMap.Enter):
			m.State = ReadFile
			m.FileID = m.List.SelectedItem().(filelist.FileItem).ID

		// Toggle help.
		case key.Matches(*msg, constants.HelpKeyMap.Help):
			m.Help.ShowAll = !m.Help.ShowAll

		// Default case.
		default:
			(*m).List, *cmd = m.List.Update(*msg)

		}

	// Case when the state is reading the file.
	case ReadFile:
		switch {

		// Quit reading the file.
		case key.Matches(*msg, constants.HelpKeyMap.Quit):
			m.State = ReadFileList

		// Toggle help.
		case key.Matches(*msg, constants.HelpKeyMap.Help):
			m.Help.ShowAll = !m.Help.ShowAll

		// Default update of a viewPort.
		default:
			(*m).ViewPort, *cmd = m.ViewPort.Update(*msg)
		}

	// Case when the file is being written.
	case WriteFile:
		switch {

		// Quit:
		case key.Matches(*msg, constants.HelpKeyMap.Quit):
			(*m).State = ReadFileList

		// Default:
		default:
			var form tea.Model
			form, *cmd = (*m).Form.Update(*msg)

			if f, ok := form.(*huh.Form); ok {
				(*m).Form = *f
			}

			//Check if the state of the form is complete:
			if (*m).Form.State == huh.StateCompleted {
				// TODO: write the current file. 
				(*m).State = ReadFileList
			}

		}
	}
}
