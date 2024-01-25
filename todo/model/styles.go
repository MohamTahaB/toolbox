package model

import "github.com/charmbracelet/lipgloss"

var SelectedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#4CAF50"))

var CursorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("201"))

var StrikeThroughStyle = lipgloss.NewStyle().
	Strikethrough(true).
	Faint(true)
