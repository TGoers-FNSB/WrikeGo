package wrikeresponse

import (
	"encoding/json"
)

func ApprovalsFromJSON(data []byte) (Approvals, error) {
	var item Approvals
	err := json.Unmarshal(data, &item)
	return item, err
}

type Approvals struct {
	Kind string `json:"kind"`
	Data []struct {
		TaskId      *string `json:"taskId,omitempty"`
		AuthorId    *string `json:"authorId,omitempty"`
		Title       *string `json:"title,omitempty"`
		Description *string `json:"description,omitempty"`
		UpdatedDate *string `json:"updatedDate,omitempty"`
		DueDate     *string `json:"dueDate,omitempty"`
		Decisions   *[]struct {
			ApproverId  *string `json:"approverId,omitempty"`
			Comment     *string `json:"comment,omitempty"`
			Status      *string `json:"status,omitempty"`
			UpdatedDate *string `json:"updatedDate,omitempty"`
		} `json:"decisions,omitempty"`
		ReviewId            *string   `json:"reviewId,omitempty"`
		AttachmentIds       *[]string `json:"attachmentIds,omitempty"`
		Type                *string   `json:"type,omitempty"`
		AutoFinishOnApprove *bool     `json:"autoFinishOnApprove,omitempty"`
		AutoFinishOnReject  *bool     `json:"autoFinishOnReject,omitempty"`
		Finished            *bool     `json:"finished,omitempty"`
		FinisherId          *string   `json:"finisherId,omitempty"`
		Id                  *string   `json:"id,omitempty"`
		Status              *string   `json:"status,omitempty"`
		FolderId            *string   `json:"folderId,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}
