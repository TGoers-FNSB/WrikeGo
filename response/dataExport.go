package wrikeresponse

import (
	"encoding/json"
)

func DataExportFromJSON(data []byte) (DataExport, error) {
	var item DataExport
	err := json.Unmarshal(data, &item)
	return item, err
}

type DataExport struct {
	Kind string `json:"kind"`
	Data []struct {
		Id            *string `json:"id,omitempty"`
		CompletedDate *string `json:"completedDate,omitempty"`
		Status        *string `json:"status,omitempty"`
		Resources     *[]struct {
			Name *string `json:"name,omitempty"`
			URL  *string `json:"url,omitempty"`
		} `json:"resources,omitempty"`
		Alias   *string `json:"alias,omitempty"`
		Columns *[]struct {
			Id         *string `json:"id,omitempty"`
			Alias      *string `json:"alias,omitempty"`
			DataType   *string `json:"dataType,omitempty"`
			ForeignKey *struct {
				TableId  *string `json:"tableId,omitempty"`
				ColumnId *string `json:"columnId,omitempty"`
			} `json:"foreignKey,omitempty"`
		} `json:"columns,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}
