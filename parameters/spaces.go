package wrikeparams

type QuerySpaces struct {
	WithArchived *bool     `url:"withArchived,omitempty"`
	UserIsMember *bool     `url:"userIsMember,omitempty"`
	Fields       *[]string `url:"fields,omitempty"`
}

type CreateSpaces struct {
	AccessType  string  `url:"accessType"`
	Title       string  `url:"title"`
	Description *string `url:"description,omitempty"`
	Members     *[]struct {
		Id           string `url:"id"`
		AccessRoleId string `url:"accessRoleId"`
		IsManager    bool   `url:"isManager"`
	} `url:"members,omitempty"`
	GuestRoleId                *string   `url:"guestRoleId,omitempty"`
	DefaultProjectWorkflowId   *string   `url:"defaultProjectWorkflowId,omitempty"`
	SuggestedProjectWorkflowId *[]string `url:"suggestedProjectWorkflowId,omitempty"`
	DefaultTaskWorkflowId      *string   `url:"defaultTaskWorkflowId,omitempty"`
	SuggestedTaskWorkflowId    *[]string `url:"suggestedTaskWorkflowId,omitempty"`
	Fields                     *[]string `url:"fields,omitempty"`
}

type ModifySpaces struct {
	AccessType  *string `url:"accessType"`
	Title       *string `url:"title"`
	Description *string `url:"description,omitempty"`
	MembersAdd  *[]struct {
		Id           string `url:"id"`
		AccessRoleId string `url:"accessRoleId"`
		IsManager    bool   `url:"isManager"`
	} `url:"membersAdd,omitempty"`
	MembersUpdate *[]struct {
		Id           string `url:"id"`
		AccessRoleId string `url:"accessRoleId"`
		IsManager    bool   `url:"isManager"`
	} `url:"membersUpdate,omitempty"`
	MembersRemove                   *[]string `url:"membersRemove,omitempty"`
	GuestRoleId                     *string   `url:"guestRoleId,omitempty"`
	DefaultProjectWorkflowId        *string   `url:"defaultProjectWorkflowId,omitempty"`
	SuggestedProjectWorkflowsAdd    *[]string `url:"suggestedProjectWorkflowsAdd,omitempty"`
	SuggestedProjectWorkflowsRemove *[]string `url:"suggestedProjectWorkflowsRemove,omitempty"`
	DefaultTaskWorkflowId           *string   `url:"defaultTaskWorkflowId,omitempty"`
	SuggestedTaskWorkflowsAdd       *[]string `url:"suggestedTaskWorkflowsAdd,omitempty"`
	SuggestedTaskWorkflowsRemove    *[]string `url:"suggestedTaskWorkflowsRemove,omitempty"`
	Fields                          *[]string `url:"fields,omitempty"`
}
