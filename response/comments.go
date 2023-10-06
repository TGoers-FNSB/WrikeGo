package wrikeresponse

import "encoding/json"

func CommentsFromJSON(data []byte) (Comments, error) {
	var item Comments
	err := json.Unmarshal(data, &item)
	return item, err
}

type Comments struct {
	Kind string `json:"kind"`
	Data []struct {
		Id                *string `json:"id,omitempty"`
		AuthorId          *string `json:"authorId,omitempty"`
		Text              *string `json:"text,omitempty"`
		UpdatedDate       *string `json:"updatedDate,omitempty"`
		CreatedDate       *string `json:"createdDate,omitempty"`
		TaskId            *string `json:"taskId,omitempty,omitempty"`
		FolderId          *string `json:"folderId,omitempty,omitempty"`
		Type              *string `json:"type,omitempty"`
		EmailSubject      *string `json:"emailSubject,omitempty"`
		Direction         *string `json:"direction,omitempty"`
		ExternalRequester *struct {
			Id        *string `json:"id,omitempty"`
			FirstName *string `json:"firstName,omitempty"`
			LastName  *string `json:"lastName,omitempty"`
			Email     *string `json:"email,omitempty"`
		} `json:"externalRequester,omitempty"`
	} `json:"data"`
}
