package wrikeresponse

import (
	"encoding/json"
)

func AccessRolesFromJSON(data []byte) (AccessRoles, error) {
	var item AccessRoles
	err := json.Unmarshal(data, &item)
	return item, err
}

type AccessRoles struct {
	Kind string `json:"kind"`
	Data []struct {
		Id          *string `json:"id,omitempty"`
		Title       *string `json:"title,omitempty"`
		Description *string `json:"description,omitempty"`
	} `json:"data"`
}