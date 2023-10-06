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
		Id            string `json:"id"`
		CompletedDate string `json:"completedDate"`
		Status        string `json:"status"`
		Resources     []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"resources"`
		Alias   string `json:"alias"`
		Columns []struct {
			Id         string `json:"id"`
			Alias      string `json:"alias"`
			DataType   string `json:"dataType"`
			ForeignKey struct {
				TableId  string `json:"tableId"`
				ColumnId string `json:"columnId"`
			} `json:"foreignKey"`
		} `json:"columns"`
	} `json:"data"`
}
