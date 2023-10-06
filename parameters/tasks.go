package wrikeparams

type QueryTasks struct {
	Descendants             *bool        `url:"descendants,omitempty"`
	Title                   *string      `url:"title,omitempty"`
	Status                  *[]string    `url:"status,omitempty"`
	Importance              *string      `url:"importance,omitempty"`
	StartDate               *DateOrRange `url:"startDate,omitempty"`
	DueDate                 *DateOrRange `url:"dueDate,omitempty"`
	ScheduledDate           *DateOrRange `url:"scheduledDate,omitempty"`
	CreatedDate             *DateOrRange `url:"createdDate,omitempty"`
	UpdatedDate             *DateOrRange `url:"updatedDate,omitempty"`
	CompletedDate           *DateOrRange `url:"completedDate,omitempty"`
	Authors                 *[]string    `url:"authors,omitempty"`
	Responsibles            *[]string    `url:"responsibles,omitempty"`
	ResponsiblePlaceholders *[]string    `url:"responsiblePlaceholders,omitempty"`
	Permalink               *string      `url:"permalink,omitempty"`
	Type                    *string      `url:"type,omitempty"`
	Limit                   *int         `url:"limit,omitempty"`
	SortField               *string      `url:"sortField,omitempty"`
	SortOrder               *string      `url:"sortOrder,omitempty"`
	SubTasks                *bool        `url:"subTasks,omitempty"`
	PageSize                *int         `url:"pageSize,omitempty"`
	NextPageToken           *string      `url:"nextPageToken,omitempty"`
	Metadata                *Metadata    `url:"metadata,omitempty"`
	CustomField             *CustomField `url:"customField,omitempty"`
	CustomStatuses          *[]string    `url:"customStatuses,omitempty"`
	BillingTypes            *[]string    `url:"billingTypes,omitempty"`
	Fields                  *[]string    `url:"fields,omitempty"`
}

type QueryTasksFieldsHistory struct {
	UpdatedDate *DateOrRange `url:"updatedDate,omitempty"`
	Fields      *[]string    `url:"fields,omitempty"`
}

type CreateTasks struct {
	Title                   string            `url:"title"`
	Description             *string           `url:"description,omitempty"`
	Status                  *string           `url:"status,omitempty"`
	Importance              *string           `url:"importance,omitempty"`
	Dates                   *TaskDates        `url:"dates,omitempty"`
	Shareds                 *[]string         `url:"shareds,omitempty"`
	Parents                 *[]string         `url:"parents,omitempty"`
	Responsibles            *[]string         `url:"responsibles,omitempty"`
	ResponsiblePlaceholders *[]string         `url:"resopnsiblePlaceholders,omitempty"`
	Followers               *[]string         `url:"followers,omitempty"`
	Follow                  *bool             `url:"follow,omitempty"`
	PriorityBefore          *string           `url:"priorityBefore,omitempty"`
	PriorityAfter           *string           `url:"priorityAfter,omitempty"`
	SuperTasks              *[]string         `url:"superTasks,omitempty"`
	Metadata                *[]Metadata       `url:"metadata,omitempty"`
	CustomFields            *[]CustomField    `url:"customFields,omitempty"`
	CustomStatus            *string           `url:"customStatus,omitempty"`
	EffortAllocation        *EffortAllocation `url:"effortAllocation,omitempty"`
	BillingType             *string           `url:"billingType,omitempty"`
	Fields                  *[]string         `url:"fields,omitempty"`
}

type ModifyTasks struct {
	Title                         *string           `url:"title,omitempty"`
	Description                   *string           `url:"description,omitempty"`
	Status                        *string           `url:"status,omitempty"`
	Importance                    *string           `url:"importance,omitempty"`
	Dates                         *TaskDates        `url:"dates,omitempty"`
	AddParents                    *[]string         `url:"addParents,omitempty"`
	RemoveParents                 *[]string         `url:"removeParents,omitempty"`
	AddShareds                    *[]string         `url:"addShareds,omitempty"`
	RemoveShareds                 *[]string         `url:"removeShareds,omitempty"`
	AddResponsibles               *[]string         `url:"addResponsibles,omitempty"`
	RemoveResponsibles            *[]string         `url:"removeResponsibles,omitempty"`
	AddResponsiblePlaceholders    *[]string         `url:"addResopnsiblePlaceholders,omitempty"`
	RemoveResponsiblePlaceholders *[]string         `url:"removeResopnsiblePlaceholders,omitempty"`
	AddFollowers                  *[]string         `url:"addFollowers,omitempty"`
	Follow                        *bool             `url:"follow,omitempty"`
	PriorityBefore                *string           `url:"priorityBefore,omitempty"`
	PriorityAfter                 *string           `url:"priorityAfter,omitempty"`
	AddSuperTasks                 *[]string         `url:"addSuperTasks,omitempty"`
	RemoveSuperTasks              *[]string         `url:"removeSuperTasks,omitempty"`
	Metadata                      *[]Metadata       `url:"metadata,omitempty"`
	CustomFields                  *[]CustomField    `url:"customFields,omitempty"`
	CustomStatus                  *string           `url:"customStatus,omitempty"`
	Restore                       *bool             `url:"restore,omitempty"`
	EffortAllocation              *EffortAllocation `url:"effortAllocation,omitempty"`
	BillingType                   *string           `url:"billingType,omitempty"`
	Fields                        *[]string         `url:"fields,omitempty"`
}