package wrikeresponse

import "encoding/json"

func ColorsFromJSON(data []byte) (Colors, error) {
	var item Colors
	err := json.Unmarshal(data, &item)
	return item, err
}

type Colors struct {
	Kind string `json:"kind"`
	Data []struct {
		Name *string `json:"name,omitempty"`
		Hex  *string `json:"hex,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}