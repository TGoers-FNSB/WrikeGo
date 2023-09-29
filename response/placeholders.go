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
		Id         string `json:"id"`
		Title      string `json:"title"`
		ShortTitle string `json:"shortTitle"`
		AvatarURL  string `json:"avatarUrl"`
		JobRoleId  string `json:"jobRoleId"`
		Deleted    bool   `json:"deleted"`
	} `json:"data"`
}
