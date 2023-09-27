package wrikeresponse

import "encoding/json"

func IDsFromJSON(data []byte) (IDs, error) {
	var item IDs
	err := json.Unmarshal(data, &item)
	return item, err
}

type IDs struct {
	Kind string `json:"kind"`
	Data []struct {
		ID      string `json:"id"`
		APIV2ID string `json:"apiV2Id"`
	} `json:"data"`
}