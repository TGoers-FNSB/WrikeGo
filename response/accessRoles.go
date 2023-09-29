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
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"data"`
}