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
		Id     string `json:"id"`
		Name   string `json:"name"`
		Order  int    `json:"order"`
		Hidden bool   `json:"hidden"`
	} `json:"data"`
}