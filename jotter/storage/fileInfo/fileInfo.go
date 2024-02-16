package fileinfo

type FileInfo struct {
	Title   string
	Content string
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
