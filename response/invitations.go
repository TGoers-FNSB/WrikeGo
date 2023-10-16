package wrikeresponse

import "encoding/json"

func InvitationsFromJSON(data []byte) (Invitations, error) {
	var item Invitations
	err := json.Unmarshal(data, &item)
	return item, err
}

type Invitations struct {
	Kind string `json:"kind"`
	Data []struct {
		Id             *string `json:"id,omitempty"`
		AccountId      *string `json:"accountId,omitempty"`
		FirstName      *string `json:"firstName,omitempty"`
		LastName       *string `json:"lastName,omitempty"`
		Email          *string `json:"email,omitempty"`
		Status         *string `json:"status,omitempty"`
		InviterUserId  *string `json:"inviterUserId,omitempty"`
		InvitationDate *string `json:"invitationDate,omitempty"`
		ResolvedDate   *string `json:"resolvedDate,omitempty"`
		Role           *string `json:"role,omitempty"`
		External       *bool   `json:"external,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}
