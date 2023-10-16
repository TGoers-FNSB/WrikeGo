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
		Id          *string `json:"id,omitempty"`
		Title       *string `json:"title,omitempty"`
		Description *string `json:"description,omitempty"`
		AvatarURL   *string `json:"avatarUrl,omitempty"`
		AccessType  *string `json:"accessType,omitempty"`
		Archived    *bool   `json:"archived,omitempty"`
		Members     *[]struct {
			Id           *string `json:"id,omitempty"`
			AccessRoleId *string `json:"accessRoleId,omitempty"`
			IsManager    *bool   `json:"isManager,omitempty"`
		} `json:"members,omitempty"`
		GuestRoleId              *string `json:"guestRoleId,omitempty"`
		DefaultProjectWorkflowId *string `json:"defaultProjectWorkflowId,omitempty"`
		DefaultTaskWorkflowId    *string `json:"defaultTaskWorkflowId,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}