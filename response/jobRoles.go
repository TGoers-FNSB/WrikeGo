package wrikeresponse

import (
	"encoding/json"
)

func JobRolesFromJSON(data []byte) (JobRoles, error) {
	var item JobRoles
	err := json.Unmarshal(data, &item)
	return item, err
}

type JobRoles struct {
	Kind string `json:"kind"`
	Data []struct {
		Id         *string `json:"id,omitempty"`
		Title      *string `json:"title,omitempty"`
		ShortTitle *string `json:"shortTitle,omitempty"`
		AvatarURL  *string `json:"avatarUrl,omitempty"`
		IsDeleted  *bool   `json:"isDeleted,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}