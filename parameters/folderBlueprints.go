package wrikeparams

type CreateFolderBlueprintsAsync struct {
	Parent             string  `url:"parent"`
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
