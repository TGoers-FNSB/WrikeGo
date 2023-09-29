package wrikeparams

type QueryApprovals struct {
	Statuses         *string      `url:"statuses,omitempty"`
	UpdatedDate      *DateOrRange `url:"updatedDate,omitempty"`
	Approvers        *[]string    `url:"approvers,omitempty"`
	PendingApprovers *[]string    `url:"pendingApprovers,omitempty"`
	Limit            *int         `url:"limit,omitempty"`
	PageSize         *int         `url:"pageSize,omitempty"`
	NextPageToken    *string      `url:"nextPageToken,omitempty"`
}

type CreateApprovals struct {
	Description         *string   `url:"description,omitempty"`
	DueDate             *string   `url:"dueDate,omitempty"`
	Approvers           *[]string `url:"approvers,omitempty"`
	Attachments         *[]string `url:"attachments,omitempty"`
	AutoFinishOnApprove *string   `url:"autoFinishOnApprove,omitempty"`
	AutoFinishOnReject  *string   `url:"autoFinishOnReject,omitempty"`
}

type ModifyApprovals struct {
	Description         *string   `url:"description,omitempty"`
	DueDate             *string   `url:"dueDate,omitempty"`
	AddApprovers        *[]string `url:"addApprovers,omitempty"`
	RemoveApprovers     *[]string `url:"removeApprovers,omitempty"`
	AddAttachments      *[]string `url:"addAttachments,omitempty"`
	RemoveAttachments   *[]string `url:"removeAttachments,omitempty"`
	AutoFinishOnApprove *string   `url:"autoFinishOnApprove,omitempty"`
	AutoFinishOnReject  *string   `url:"autoFinishOnReject,omitempty"`
}
