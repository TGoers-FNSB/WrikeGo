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
		Id              *string      `json:"id,omitempty"`
		Status          *string      `json:"status,omitempty"`
		ProgressPercent *int         `json:"progressPercent,omitempty"`
		TotalCount      *int         `json:"totalCount,omitempty"`
		ProcessedCount  *int         `json:"processedCount,omitempty"`
		Type            *string      `json:"type,omitempty"`
		Result          *interface{} `json:"result,omitempty"`
		ErrorMessage    *string      `json:"errorMessage,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}
