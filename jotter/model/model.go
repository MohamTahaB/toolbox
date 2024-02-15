package model

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
)

type state int

const (
	ReadFileList state = iota
	WriteFileList
	ReadFile
	WriteFile
)

type Model struct {
	State state
	List  list.Model
	FileEdit textarea.Model
	ViewPort viewport.Model
}