package wrikeparams

type QueryApprovals struct {
	Statuses         string      `url:"statuses,omitempty"`
	UpdatedDate      DateOrRange `url:"updatedDate,omitempty,struct"`
	Approvers        []string    `url:"approvers,omitempty,slice"`
	PendingApprovers []string    `url:"pendingApprovers,omitempty,slice"`
	Limit            int         `url:"limit,omitempty"`
	PageSize         int         `url:"pageSize,omitempty"`
	NextPageToken    string      `url:"nextPageToken,omitempty"`
}

type CreateApprovals struct {
	Description         string   `url:"description,omitempty"`
	DueDate             string   `url:"dueDate,omitempty"`
	Approvers           []string `url:"approvers,omitempty,slice"`
	Attachments         []string `url:"attachments,omitempty,slice"`
	AutoFinishOnApprove string   `url:"autoFinishOnApprove,omitempty"`
	AutoFinishOnReject  string   `url:"autoFinishOnReject,omitempty"`
}

type ModifyApprovals struct {
	Description         string   `url:"description,omitempty"`
	DueDate             string   `url:"dueDate,omitempty"`
	AddApprovers        []string `url:"addApprovers,omitempty,slice"`
	RemoveApprovers     []string `url:"removeApprovers,omitempty,slice"`
	AddAttachments      []string `url:"addAttachments,omitempty,slice"`
	RemoveAttachments   []string `url:"removeAttachments,omitempty,slice"`
	AutoFinishOnApprove string   `url:"autoFinishOnApprove,omitempty"`
	AutoFinishOnReject  string   `url:"autoFinishOnReject,omitempty"`
}
