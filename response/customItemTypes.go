package wrikeresponse

import (
	"encoding/json"
)

func CustomItemTypesFromJSON(data []byte) (CustomItemTypes, error) {
	var item CustomItemTypes
	err := json.Unmarshal(data, &item)
	return item, err
}

type CustomItemTypes struct {
	Kind string `json:"kind"`
	Data []struct {
		Id          string `json:"id"`
		Title       string `json:"title"`
		RelatedType string `json:"relatedType"`
		SpaceID     string `json:"spaceId"`
		Description string `json:"description"`
	} `json:"data"`
}