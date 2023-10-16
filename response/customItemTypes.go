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
		Id          *string `json:"id,omitempty"`
		Title       *string `json:"title,omitempty"`
		RelatedType *string `json:"relatedType,omitempty"`
		SpaceId     *string `json:"spaceId,omitempty"`
		Description *string `json:"description,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}