package wrikeparams

type QueryTaskBlueprints struct {
	Limit         *int    `url:"limit,omitempty"`
	PageSize      *int    `url:"pageSize,omitempty"`
	NextPageToken *string `url:"nextPageToken,omitempty"`
}

type CreateTaskBlueprintsAsync struct {
	SuperTaskId        *string `url:"superTaskId,omitempty"`
	ParentId           *string `url:"parentId,omitempty"`
	Title              string  `url:"title"`
	TitlePrefix        *string `url:"titlePrefix,omitempty"`
	CopyDescriptions   *bool   `url:"copyDescriptions,omitempty"`
	NotifyResponsibles *bool   `url:"notifyResponsibles,omitempty"`
	CopyResponsibles   *bool   `url:"copyResponsibles,omitempty"`
	CopyCustomFields   *bool   `url:"copyCustomFields,omitempty"`
	CopyAttachments    *bool   `url:"copyAttachments,omitempty"`
	RescheduleDate     *string `url:"rescheduleDate,omitempty"`
	RescheduleMode     *string `url:"rescheduleMode,omitempty"`
	EntryLimit         *int    `url:"entryLimit,omitempty"`
}
