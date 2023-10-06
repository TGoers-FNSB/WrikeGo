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
		TaskId      string `json:"taskId,omitempty"`
		AuthorId    string `json:"authorId"`
		Title       string `json:"title"`
		Description string `json:"description"`
		UpdatedDate string `json:"updatedDate"`
		DueDate     string `json:"dueDate"`
		Decisions   []struct {
			ApproverId  string `json:"approverId"`
			Comment     string `json:"comment"`
			Status      string `json:"status"`
			UpdatedDate string `json:"updatedDate"`
		} `json:"decisions"`
		ReviewId            string   `json:"reviewId"`
		AttachmentIds       []string `json:"attachmentIds"`
		Type                string   `json:"type"`
		AutoFinishOnApprove bool     `json:"autoFinishOnApprove"`
		AutoFinishOnReject  bool     `json:"autoFinishOnReject"`
		Finished            bool     `json:"finished"`
		FinisherId          string   `json:"finisherId"`
		Id                  string   `json:"id"`
		Status              string   `json:"status"`
		FolderId            string   `json:"folderId,omitempty"`
	} `json:"data"`
}
