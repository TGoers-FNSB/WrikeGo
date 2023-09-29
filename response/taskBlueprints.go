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
		Id              string      `json:"id"`
		Title           string      `json:"title"`
		ChildIds        []string    `json:"childIds"`
		Scope           string      `json:"scope"`
		Status          string      `json:"status"`
		ProgressPercent int         `json:"progressPercent"`
		TotalCount      int         `json:"totalCount"`
		ProcessedCount  int         `json:"processedCount"`
		Type            string      `json:"type"`
		Result          interface{} `json:"result"`
		ErrorMessage    string      `json:"errorMessage"`
	} `json:"data"`
}
