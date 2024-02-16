package filelist

type FileDesignation struct {
	FileTitle string `json:"title"`
	FileDesc  string `json:"description"`
}

func (i FileDesignation) Title() string       { return i.FileTitle }
func (i FileDesignation) Description() string { return i.FileDesc }
func (i FileDesignation) FilterValue() string { return i.FileTitle }
