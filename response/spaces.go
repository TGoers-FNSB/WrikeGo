package wrikeresponse

import (
	"encoding/json"
)

func SpacesFromJSON(data []byte) (Spaces, error) {
	var item Spaces
	err := json.Unmarshal(data, &item)
	return item, err
}

type Spaces struct {
	Kind string `json:"kind"`
	Data []struct {
		Id          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description,omitempty"`
		AvatarURL   string `json:"avatarUrl"`
		AccessType  string `json:"accessType"`
		Archived    bool   `json:"archived"`
		Members     []struct {
			Id           string `json:"id"`
			AccessRoleId string `json:"accessRoleId"`
			IsManager    bool   `json:"isManager"`
		} `json:"members"`
		GuestRoleId              string `json:"guestRoleId,omitempty"`
		DefaultProjectWorkflowId string `json:"defaultProjectWorkflowId"`
		DefaultTaskWorkflowId    string `json:"defaultTaskWorkflowId"`
	} `json:"data"`
}