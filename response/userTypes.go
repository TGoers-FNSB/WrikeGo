package wrikeresponse

import (
	"encoding/json"
)

func UserTypesFromJSON(data []byte) (UserTypes, error) {
	var item UserTypes
	err := json.Unmarshal(data, &item)
	return item, err
}

type UserTypes struct {
	Kind string `json:"kind"`
	Data []struct {
		Id          *string `json:"id,omitempty"`
		Title       *string `json:"title,omitempty"`
		Description *string `json:"description,omitempty"`
	} `json:"data"`
}