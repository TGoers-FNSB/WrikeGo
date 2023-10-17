package wrikeparams

type QueryTasks struct {
	Descendants             bool        `url:"descendants,omitempty"`
	Title                   string      `url:"title,omitempty"`
	Status                  []string    `url:"status,omitempty,slice"`
	Importance              string      `url:"importance,omitempty"`
	StartDate               DateOrRange `url:"startDate,omitempty,struct"`
	DueDate                 DateOrRange `url:"dueDate,omitempty,struct"`
	ScheduledDate           DateOrRange `url:"scheduledDate,omitempty,struct"`
	CreatedDate             DateOrRange `url:"createdDate,omitempty,struct"`
	UpdatedDate             DateOrRange `url:"updatedDate,omitempty,struct"`
	CompletedDate           DateOrRange `url:"completedDate,omitempty,struct"`
	Authors                 []string    `url:"authors,omitempty,slice"`
	Responsibles            []string    `url:"responsibles,omitempty,slice"`
	ResponsiblePlaceholders []string    `url:"responsiblePlaceholders,omitempty,slice"`
	Permalink               string      `url:"permalink,omitempty"`
	Type                    string      `url:"type,omitempty"`
	Limit                   int         `url:"limit,omitempty"`
	SortField               string      `url:"sortField,omitempty"`
	SortOrder               string      `url:"sortOrder,omitempty"`
	SubTasks                bool        `url:"subTasks,omitempty"`
	PageSize                int         `url:"pageSize,omitempty"`
	NextPageToken           string      `url:"nextPageToken,omitempty"`
	Metadata                Metadata    `url:"metadata,omitempty,struct"`
	CustomField             CustomField `url:"customField,omitempty,struct"`
	CustomStatuses          []string    `url:"customStatuses,omitempty,slice"`
	BillingTypes            []string    `url:"billingTypes,omitempty,slice"`
	Fields                  []string    `url:"fields,omitempty,slice"`
}

type QueryTasksFieldsHistory struct {
	UpdatedDate DateOrRange `url:"updatedDate,omitempty"`
	Fields      []string    `url:"fields,omitempty,slice"`
}

type CreateTasks struct {
	Title                   string           `url:"title"`
	Description             string           `url:"description,omitempty"`
	Status                  string           `url:"status,omitempty"`
	Importance              string           `url:"importance,omitempty"`
	Dates                   TaskDates        `url:"dates,omitempty,struct"`
	Shareds                 []string         `url:"shareds,omitempty,slice"`
	Parents                 []string         `url:"parents,omitempty,slice"`
	Responsibles            []string         `url:"responsibles,omitempty,slice"`
	ResponsiblePlaceholders []string         `url:"resopnsiblePlaceholders,omitempty,slice"`
	Followers               []string         `url:"followers,omitempty,slice"`
	Follow                  bool             `url:"follow,omitempty"`
	PriorityBefore          string           `url:"priorityBefore,omitempty"`
	PriorityAfter           string           `url:"priorityAfter,omitempty"`
	SuperTasks              []string         `url:"superTasks,omitempty,slice"`
	Metadata                []Metadata       `url:"metadata,omitempty,slice+struct"`
	CustomFields            []CustomField    `url:"customFields,omitempty,slice+struct"`
	CustomStatus            string           `url:"customStatus,omitempty"`
	EffortAllocation        EffortAllocation `url:"effortAllocation,omitempty,struct"`
	BillingType             string           `url:"billingType,omitempty"`
	Fields                  []string         `url:"fields,omitempty,slice"`
}

type ModifyTasks struct {
	Title                         string           `url:"title,omitempty"`
	Description                   string           `url:"description,omitempty"`
	Status                        string           `url:"status,omitempty"`
	Importance                    string           `url:"importance,omitempty"`
	Dates                         TaskDates        `url:"dates,omitempty,struct"`
	AddParents                    []string         `url:"addParents,omitempty,slice"`
	RemoveParents                 []string         `url:"removeParents,omitempty,slice"`
	AddShareds                    []string         `url:"addShareds,omitempty,slice"`
	RemoveShareds                 []string         `url:"removeShareds,omitempty,slice"`
	AddResponsibles               []string         `url:"addResponsibles,omitempty,slice"`
	RemoveResponsibles            []string         `url:"removeResponsibles,omitempty,slice"`
	AddResponsiblePlaceholders    []string         `url:"addResopnsiblePlaceholders,omitempty,slice"`
	RemoveResponsiblePlaceholders []string         `url:"removeResopnsiblePlaceholders,omitempty,slice"`
	AddFollowers                  []string         `url:"addFollowers,omitempty,slice"`
	Follow                        bool             `url:"follow,omitempty"`
	PriorityBefore                string           `url:"priorityBefore,omitempty"`
	PriorityAfter                 string           `url:"priorityAfter,omitempty"`
	AddSuperTasks                 []string         `url:"addSuperTasks,omitempty,slice"`
	RemoveSuperTasks              []string         `url:"removeSuperTasks,omitempty,slice"`
	Metadata                      []Metadata       `url:"metadata,omitempty,slice+struct"`
	CustomFields                  []CustomField    `url:"customFields,omitempty,slice+struct"`
	CustomStatus                  string           `url:"customStatus,omitempty"`
	Restore                       bool             `url:"restore,omitempty"`
	EffortAllocation              EffortAllocation `url:"effortAllocation,omitempty,struct"`
	BillingType                   string           `url:"billingType,omitempty"`
	Fields                        []string         `url:"fields,omitempty,slice"`
}
