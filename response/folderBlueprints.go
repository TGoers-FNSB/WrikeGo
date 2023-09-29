package wrikeresponse

import (
	"encoding/json"
)

func FolderBlueprintsFromJSON(data []byte) (FolderBlueprints, error) {
	var item FolderBlueprints
	err := json.Unmarshal(data, &item)
	return item, err
}

type FolderBlueprints struct {
	Kind string `json:"kind"`
	Data []struct {
		Id              string       `json:"id"`
		Status          string       `json:"status"`
		ProgressPercent int          `json:"progressPercent"`
		TotalCount      int         `json:"totalCount"`
		ProcessedCount  int         `json:"processedCount"`
		Type            string       `json:"type"`
		Result          interface{} `json:"result"`
		ErrorMessage    string      `json:"errorMessage"`
	} `json:"data"`
}
