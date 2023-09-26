package wrikeparams

type CreateWorkflows struct {
	Name string `url:"name"`
}

type ModifyWorkflows struct {
	Name         *string `url:"name,omitempty"`
	Hidden       *bool   `url:"hidden,omitempty"`
	CustomStatus *struct {
		ID           *string `url:"id,omitempty"`
		Name         *string `url:"name,omitempty"`
		StandardName *bool   `url:"standardName,omitempty"`
		Color        *string `url:"color,omitempty"`
		Standard     *bool   `url:"standard,omitempty"`
		Group        *string `url:"group,omitempty"`
		Hidden       *bool   `url:"hidden,omitempty"`
	} `url:"customStatus,omitempty"`
}
