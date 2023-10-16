package wrikeresponse

import "encoding/json"

func TimelogCategoriesFromJSON(data []byte) (TimelogCategories, error) {
	var item TimelogCategories
	err := json.Unmarshal(data, &item)
	return item, err
}

type TimelogCategories struct {
	Kind string `json:"kind"`
	Data []struct {
		Id     *string `json:"id,omitempty"`
		Name   *string `json:"name,omitempty"`
		Order  *int    `json:"order,omitempty"`
		Hidden *bool   `json:"hidden,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}