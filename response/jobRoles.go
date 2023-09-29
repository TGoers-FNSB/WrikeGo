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
		Id         string `json:"id"`
		Title      string `json:"title"`
		ShortTitle string `json:"shortTitle"`
		AvatarURL  string `json:"avatarUrl"`
		IsDeleted  bool   `json:"isDeleted"`
	} `json:"data"`
}