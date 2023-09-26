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
		Id                string `json:"id"`
		AuthorId          string `json:"authorId"`
		Text              string `json:"text"`
		UpdatedDate       string `json:"updatedDate"`
		CreatedDate       string `json:"createdDate"`
		TaskId            string `json:"taskId,omitempty"`
		FolderId          string `json:"folderId,omitempty"`
		Type              string `json:"type"`
		EmailSubject      string `json:"emailSubject"`
		Direction         string `json:"direction"`
		ExternalRequester struct {
			Id        string `json:"id"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Email     string `json:"email"`
		} `json:"externalRequester"`
	} `json:"data"`
}
