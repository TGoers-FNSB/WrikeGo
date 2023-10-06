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
		Id              *string `json:"id,omitempty"`
		Status          *string `json:"status,omitempty"`
		ProgressPercent *int    `json:"progressPercent,omitempty"`
		TotalCount      *int    `json:"totalCount,omitempty"`
		ProcessedCount  *int    `json:"processedCount,omitempty"`
		Type            *string `json:"type,omitempty"`
		Result          *struct {
			FolderId *string `json:"folderId,omitempty"`
		} `json:"result,omitempty"`
		ErrorMessage *string `json:"errorMessage,omitempty"`
	} `json:"data"`
}
