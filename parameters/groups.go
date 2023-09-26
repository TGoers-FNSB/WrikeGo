package wrikeparams

type QueryGroups struct {
	Metadata  *Metadata `url:"metadata,omitempty"`
	PageSize  *float64  `url:"pageSize,omitempty"`
	PageToken *string   `url:"pageToken,omitempty"`
	Fields    *[]string `url:"fields,omitempty"`
}

type CreateGroups struct {
	Title    string      `url:"title"`
	Members  *[]string   `url:"members,omitempty"`
	Parent   *string     `url:"parent,omitempty"`
	Avatar   *Avatar     `url:"avatar,omitempty"`
	Metadata *[]Metadata `url:"metadata,omitempty"`
}

type ModifyGroupsBulk struct {
	Members []struct {
		Id            string  `url:"id"`
		AddMembers    *string `url:"addMembers,omitempty"`
		RemoveMembers *string `url:"removeMembers,omitempty"`
	} `url:"members"`
}

type ModifyGroups struct {
	Title             *string     `url:"title,omitempty"`
	AddMembers        *[]string   `url:"addMembers,omitempty"`
	RemoveMembers     *[]string   `url:"removeMembers,omitempty"`
	AddInvitations    *[]string   `url:"addInvitations,omitempty"`
	RemoveInvitations *[]string   `url:"removeInvitations,omitempty"`
	Parent            *string     `url:"parent,omitempty"`
	Avatar            *Avatar     `url:"avatar,omitempty"`
	Metadata          *[]Metadata `url:"metadata,omitempty"`
}

type DeleteGroups struct {
	Test *bool `url:"test,omitempty"`
}
