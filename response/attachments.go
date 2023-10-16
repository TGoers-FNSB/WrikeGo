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
		Id                  *string   `json:"id,omitempty"`
		AuthorId            *string   `json:"authorId,omitempty"`
		Name                *string   `json:"name,omitempty"`
		CreatedDate         *string   `json:"createdDate,omitempty"`
		Version             *int      `json:"version,omitempty"`
		Type                *string   `json:"type,omitempty"`
		ContentType         *string   `json:"contentType,omitempty"`
		Size                *int      `json:"size,omitempty"`
		TaskId              *string   `json:"taskId,omitempty"`
		FolderId            *string   `json:"folderId,omitempty"`
		CommentId           *string   `json:"commentId,omitempty"`
		CurrentAttachmentId *string   `json:"currentAttachmentId,omitempty"`
		PreviewUrl          *string   `json:"previewUrl,omitempty"`
		Url                 *string   `json:"url,omitempty"`
		PlaylistUrl         *string   `json:"playlistUrl,omitempty"`
		ReviewIds           *[]string `json:"reviewIds,omitempty"`
		Width               *int      `json:"width,omitempty"`
		Height              *int      `json:"height,omitempty"`
		OriginVersionId     *string   `json:"originVersionId,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}
