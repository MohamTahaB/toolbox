package filelist

type FileDesignation struct {
	FileTitle string `json:"title"`
	FileDesc  string `json:"description"`
}

type FileItem struct {
	ID        string
	FileTitle string
	FileDesc  string
}

func (i FileItem) Title() string       { return i.FileTitle }
func (i FileItem) Description() string { return i.FileDesc }
func (i FileItem) FilterValue() string { return i.FileTitle }
