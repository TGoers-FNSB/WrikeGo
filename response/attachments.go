package wrikeresponse

import "encoding/json"

func AttachmentsFromJSON(data []byte) (Attachments, error) {
	var item Attachments
	err := json.Unmarshal(data, &item)
	return item, err
}

type Attachments struct {
	Kind string `json:"kind"`
	Data []struct {
		ID                  string `json:"id"`
		AuthorID            string `json:"authorId"`
		Name                string `json:"name"`
		CreatedDate         string `json:"createdDate"`
		Version             int    `json:"version"`
		Type                string `json:"type"`
		ContentType         string `json:"contentType"`
		Size                int    `json:"size"`
		TaskID              string `json:"taskId,omitempty"`
		CurrentAttachmentID string `json:"currentAttachmentId,omitempty"`
		OriginVersionID     string `json:"originVersionId"`
		FolderID            string `json:"folderId,omitempty"`
	} `json:"data"`
}
