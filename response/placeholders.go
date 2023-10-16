package wrikeresponse

import (
	"encoding/json"
)

func PlaceholdersFromJSON(data []byte) (Placeholders, error) {
	var item Placeholders
	err := json.Unmarshal(data, &item)
	return item, err
}

type Placeholders struct {
	Kind string `json:"kind"`
	Data []struct {
		Id         *string `json:"id,omitempty"`
		Title      *string `json:"title,omitempty"`
		ShortTitle *string `json:"shortTitle,omitempty"`
		AvatarURL  *string `json:"avatarUrl,omitempty"`
		JobRoleId  *string `json:"jobRoleId,omitempty"`
		Deleted    *bool   `json:"deleted,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}
