package wrikeresponse

import (
	"encoding/json"
)

func AsyncJobFromJSON(data []byte) (AsyncJob, error) {
	var item AsyncJob
	err := json.Unmarshal(data, &item)
	return item, err
}

type AsyncJob struct {
	Kind string `json:"kind"`
	Data []struct {
		Id              string `json:"id"`
		Status          string `json:"status"`
		ProgressPercent int    `json:"progressPercent"`
		TotalCount      int    `json:"totalCount"`
		ProcessedCount  int    `json:"processedCount"`
		Type            string `json:"type"`
		Result          struct {
			FolderId string `json:"folderId"`
		} `json:"result"`
		ErrorMessage string `json:"errorMessage"`
	} `json:"data"`
}
