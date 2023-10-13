package wrikeparams

type QuerySpaces struct {
	WithArchived bool     `url:"withArchived,omitempty"`
	UserIsMember bool     `url:"userIsMember,omitempty"`
	Fields       []string `url:"fields,omitempty,slice"`
}

type CreateSpaces struct {
	AccessType                 string    `url:"accessType"`
	Title                      string    `url:"title"`
	Description                string    `url:"description,omitempty"`
	Members                    []Members `url:"members,omitempty,slice+struct"`
	GuestRoleId                string    `url:"guestRoleId,omitempty"`
	DefaultProjectWorkflowId   string    `url:"defaultProjectWorkflowId,omitempty"`
	SuggestedProjectWorkflowId []string  `url:"suggestedProjectWorkflowId,omitempty,slice"`
	DefaultTaskWorkflowId      string    `url:"defaultTaskWorkflowId,omitempty"`
	SuggestedTaskWorkflowId    []string  `url:"suggestedTaskWorkflowId,omitempty,slice"`
	Fields                     []string  `url:"fields,omitempty,slice"`
}

type ModifySpaces struct {
	AccessType                      string    `url:"accessType,omitempty"`
	Title                           string    `url:"title,omitempty"`
	Description                     string    `url:"description,omitempty"`
	MembersAdd                      []Members `url:"membersAdd,omitempty,slice+struct"`
	MembersUpdate                   []Members `url:"membersUpdate,omitempty,slice+struct"`
	MembersRemove                   []string  `url:"membersRemove,omitempty,slice"`
	GuestRoleId                     string    `url:"guestRoleId,omitempty"`
	DefaultProjectWorkflowId        string    `url:"defaultProjectWorkflowId,omitempty"`
	SuggestedProjectWorkflowsAdd    []string  `url:"suggestedProjectWorkflowsAdd,omitempty,slice"`
	SuggestedProjectWorkflowsRemove []string  `url:"suggestedProjectWorkflowsRemove,omitempty,slice"`
	DefaultTaskWorkflowId           string    `url:"defaultTaskWorkflowId,omitempty"`
	SuggestedTaskWorkflowsAdd       []string  `url:"suggestedTaskWorkflowsAdd,omitempty,slice"`
	SuggestedTaskWorkflowsRemove    []string  `url:"suggestedTaskWorkflowsRemove,omitempty,slice"`
	Fields                          []string  `url:"fields,omitempty,slice"`
}
