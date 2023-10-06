package wrikeresponse

import (
	"encoding/json"
)

func TaskBlueprintsFromJSON(data []byte) (TaskBlueprints, error) {
	var item TaskBlueprints
	err := json.Unmarshal(data, &item)
	return item, err
}

type TaskBlueprints struct {
	Kind string `json:"kind"`
	Data []struct {
		Id              *string      `json:"id,omitempty"`
		Title           *string      `json:"title,omitempty"`
		ChildIds        *[]string    `json:"childIds,omitempty"`
		Scope           *string      `json:"scope,omitempty"`
		Status          *string      `json:"status,omitempty"`
		ProgressPercent *int         `json:"progressPercent,omitempty"`
		TotalCount      *int         `json:"totalCount,omitempty"`
		ProcessedCount  *int         `json:"processedCount,omitempty"`
		Type            *string      `json:"type,omitempty"`
		Result          *interface{} `json:"result,omitempty"`
		ErrorMessage    *string      `json:"errorMessage,omitempty"`
	} `json:"data"`
}
