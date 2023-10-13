package wrikeparams

type QueryFolders struct {
	Permalink     string      `url:"permalink,omitempty"`
	Descendants   bool        `url:"descendants,omitempty"`
	Metadata      Metadata    `url:"metadata,omitempty,struct"`
	CustomField   CustomField `url:"customField,omitempty,struct"`
	UpdatedDate   DateOrRange `url:"updatedDate,omitempty,struct"`
	Project       bool        `url:"project,omitempty"`
	Deleted       bool        `url:"deleted,omitempty"`
	ContractTypes []string    `url:"contractTypes,omitempty,slice"`
	Fields        []string    `url:"fields,omitempty,slice"`
}

type CreateFolders struct {
	Title         string        `url:"title"`
	Description   string        `url:"description,omitempty"`
	Shareds       []string      `url:"shareds,omitempty,slice"`
	Metadata      []Metadata    `url:"metadata,omitempty,slice+struct"`
	CustomFields  []CustomField `url:"customFields,omitempty,slice+struct"`
	CustomColumns []string      `url:"customColumns,omitempty,slice"`
	Project       Project       `url:"project,omitempty,struct,bypass"`
	Fields        []string      `url:"fields,omitempty,slice"`
}

type CreateFoldersCopy struct {
	Parent             string   `url:"parent"`
	Title              string   `url:"title"`
	TitlePrefix        string   `url:"titlePrefix,omitempty"`
	CopyDescriptions   bool     `url:"copyDescriptions,omitempty"`
	CopyResponsibles   bool     `url:"copyResponsibles,omitempty"`
	AddResponsibles    []string `url:"addResponsibles,omitempty,slice"`
	RemoveResponsibles []string `url:"removeResponsibles,omitempty,slice"`
	CopyCustomFields   bool     `url:"copyCustomFields,omitempty"`
	CopyCustomStatuses bool     `url:"copyCustomStatuses,omitempty"`
	CopyStatuses       bool     `url:"copyStatuses,omitempty"`
	CopyParents        bool     `url:"copyParents,omitempty"`
	RescheduleDate     string   `url:"rescheduleDate,omitempty"`
	RescheduleMode     string   `url:"rescheduleMode,omitempty"`
	EntryLimit         int      `url:"entryLimit,omitempty"`
	AddShareds         []string `url:"addShareds,omitempty,slice"`
	RemoveShareds      []string `url:"removeShareds,omitempty,slice"`
}

type ModifyFolders struct {
	Title              string        `url:"title,omitempty"`
	Description        string        `url:"description,omitempty"`
	AddParents         []string      `url:"addParents,omitempty,slice"`
	RemoveParents      []string      `url:"removeParents,omitempty,slice"`
	AddShareds         []string      `url:"addShareds,omitempty,slice"`
	RemoveShareds      []string      `url:"removeShareds,omitempty,slice"`
	Metadata           []Metadata    `url:"metadata,omitempty,slice+struct"`
	Restore            bool          `url:"restore,omitempty"`
	CustomFields       []CustomField `url:"customFields,omitempty,slice+struct"`
	CustomColumns      []string      `url:"customColumns,omitempty,slice"`
	ClearCustomColumns bool          `url:"clearCustomColumns,omitempty"`
	Project            Project       `url:"project,omitempty,struct"`
	Fields             []string      `url:"fields,omitempty,slice"`
}
