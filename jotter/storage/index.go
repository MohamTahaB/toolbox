package storage

import (
	"fmt"
	filelist "toolbox/jotter/storage/fileList"
)

func GetListDesignation(JSONDir, id string) (*filelist.FileDesignation, error) {

	// Start by unmarshaling the json file.
	listsMap, err := unmarshal(JSONDir)
	if err != nil {
		return nil, err
	}

	if _, ok := (*listsMap)[id]; !ok {
		return nil, fmt.Errorf("error retrieving the .md file of id %s: file does not exist in the repertory", id)
	}

	// Return the list designation.
	ld := (*listsMap)[id]
	return &ld, nil

}
