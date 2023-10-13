package wrikeparams

type QueryGroups struct {
	Metadata  Metadata `url:"metadata,omitempty,struct"`
	PageSize  int      `url:"pageSize,omitempty"`
	PageToken string   `url:"pageToken,omitempty"`
	Fields    []string `url:"fields,omitempty,slice"`
}

type CreateGroups struct {
	Title    string     `url:"title"`
	Members  []string   `url:"members,omitempty,slice"`
	Parent   string     `url:"parent,omitempty"`
	Avatar   Avatar     `url:"avatar,omitempty,struct"`
	Metadata []Metadata `url:"metadata,omitempty,slice+struct"`
}

type ModifyGroupsBulk struct {
	Members []struct {
		Id            string `url:"id"`
		AddMembers    string `url:"addMembers,omitempty"`
		RemoveMembers string `url:"removeMembers,omitempty"`
	} `url:"members,slice+struct"`
}

type ModifyGroups struct {
	Title             string     `url:"title,omitempty"`
	AddMembers        []string   `url:"addMembers,omitempty,slice"`
	RemoveMembers     []string   `url:"removeMembers,omitempty,slice"`
	AddInvitations    []string   `url:"addInvitations,omitempty,slice"`
	RemoveInvitations []string   `url:"removeInvitations,omitempty,slice"`
	Parent            string     `url:"parent,omitempty"`
	Avatar            Avatar     `url:"avatar,omitempty,struct"`
	Metadata          []Metadata `url:"metadata,omitempty,slice+struct"`
}

type DeleteGroups struct {
	Test bool `url:"test,omitempty"`
}
