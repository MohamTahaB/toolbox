package storage

import (
	"encoding/json"
	"fmt"
	"os"
	filelist "toolbox/jotter/storage/fileList"
)

// unmarshals the json file into a map of key = id, and value = name.
func unmarshal(JSONDir string) (*map[string]filelist.FileDesignation, error) {
	// Read the json file.
	b, err := os.ReadFile(JSONDir)
	if err != nil {
		return nil, fmt.Errorf("error reading the json file: %v", err)
	}

	// Unmarshal the content of the json file.
	var IDmap map[string]filelist.FileDesignation

	if err = json.Unmarshal(b, &IDmap); err != nil {
		return nil, fmt.Errorf("error unmarshaling the json file: %v", err)
	}

	return &IDmap, nil
}

// marshals the id map.
func marshal(IDmap *map[string]filelist.FileDesignation) (*[]byte, error) {

	b, err := json.Marshal(*IDmap)
	if err != nil {
		return nil, fmt.Errorf("error marshaling the id map: %v", err)
	}

	return &b, nil
}

// Overwrites the content of the json file with the byte slice passed as a parameter.
func editJSON(JSONDir string, content *[]byte) error {
	if err := os.WriteFile(JSONDir, *content, 0755); err != nil {
		return fmt.Errorf("error writing into the json file: %v", err)
	}
	return nil
}
