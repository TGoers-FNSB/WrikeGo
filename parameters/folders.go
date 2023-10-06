package wrikeparams

type QueryFolders struct {
	Permalink     *string      `url:"permalink,omitempty"`
	Descendants   *bool        `url:"descendants,omitempty"`
	Metadata      *Metadata    `url:"metadata,omitempty"`
	CustomField   *CustomField `url:"customField,omitempty"`
	UpdatedDate   *DateOrRange `url:"updatedDate,omitempty"`
	Project       *bool        `url:"project,omitempty"`
	Deleted       *bool        `url:"deleted,omitempty"`
	ContractTypes *[]string    `url:"contractTypes,omitempty"`
	Fields        *[]string    `url:"fields,omitempty"`
}

type CreateFolders struct {
	Title         string         `url:"title"`
	Description   *string        `url:"description,omitempty"`
	Shareds       *[]string      `url:"shareds,omitempty"`
	Metadata      *[]Metadata    `url:"metadata,omitempty"`
	CustomFields  *[]CustomField `url:"customFields,omitempty"`
	CustomColumns *[]string      `url:"customColumns,omitempty"`
	Project       *Project       `url:"project,omitempty"`
	Fields        *[]string      `url:"fields,omitempty"`
}

type CreateFoldersCopy struct {
	Parent             string    `url:"parent"`
	Title              string    `url:"title"`
	TitlePrefix        *string   `url:"titlePrefix,omitempty"`
	CopyDescriptions   *bool     `url:"copyDescriptions,omitempty"`
	CopyResponsibles   *bool     `url:"copyResponsibles,omitempty"`
	AddResponsibles    *[]string `url:"addResponsibles,omitempty"`
	RemoveResponsibles *[]string `url:"removeResponsibles,omitempty"`
	CopyCustomFields   *bool     `url:"copyCustomFields,omitempty"`
	CopyCustomStatuses *bool     `url:"copyCustomStatuses,omitempty"`
	CopyStatuses       *bool     `url:"copyStatuses,omitempty"`
	CopyParents        *bool     `url:"copyParents,omitempty"`
	RescheduleDate     *string   `url:"rescheduleDate,omitempty"`
	RescheduleMode     *string   `url:"rescheduleMode,omitempty"`
	EntryLimit         *int      `url:"entryLimit,omitempty"`
	AddShareds         *[]string `url:"addShareds,omitempty"`
	RemoveShareds      *[]string `url:"removeShareds,omitempty"`
}

type ModifyFolders struct {
	Title              *string        `url:"title,omitempty"`
	Description        *string        `url:"description,omitempty"`
	AddParents         *[]string      `url:"addParents,omitempty"`
	RemoveParents      *[]string      `url:"removeParents,omitempty"`
	AddShareds         *[]string      `url:"addShareds,omitempty"`
	RemoveShareds      *[]string      `url:"removeShareds,omitempty"`
	Metadata           *[]Metadata    `url:"metadata,omitempty"`
	Restore            *bool          `url:"restore,omitempty"`
	CustomFields       *[]CustomField `url:"customFields,omitempty"`
	CustomColumns      *[]string      `url:"customColumns,omitempty"`
	ClearCustomColumns *bool          `url:"clearCustomColumns,omitempty"`
	Project            *Project       `url:"project,omitempty"`
	Fields             *[]string      `url:"fields,omitempty"`
}
