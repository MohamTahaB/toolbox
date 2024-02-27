package model

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Msg, tea.Cmd) {
	switch msg := msg.(type) {
	// The sole change is the size of the window.
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil

	// Case when the msg is a key msg.
	case tea.KeyMsg:
		handleKeyMsg(&m, &msg)
	}

	return nil, nil
}

func handleKeyMsg(m *Model, msg *tea.KeyMsg) {
	switch m.State {

	// Case when the state is reading the file list.
	case ReadFileList:

	}
}
