package fileinfo

import "github.com/google/uuid"

type FileInfo struct {
	ID      string
	Title   string
	Content string
}

// Creates a new instance of file info.
func NewFile() *FileInfo {
	return &FileInfo{
		ID: uuid.NewString(),
		Title: "",
        Content: "",
	}
}

func (fi *FileInfo) GetTitle() string {
	if fi == nil {
		return ""
	}
	return fi.Title
}

func (fi *FileInfo) GetContent() string {
	if fi == nil {
		return ""
	}
	return fi.Content
}

func (fi *FileInfo) InitiateFile() {
	fi.Title = ""
	fi.Content = ""
}
