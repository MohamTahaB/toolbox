package storage

type StorageFile struct {
	Path string
}

// Getters.
func (sf *StorageFile) GetPath() string {
	if (sf == nil) {
		return ""
	}
	return sf.Path
}
