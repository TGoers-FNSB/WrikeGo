package wrikeparams

type CreateCustomFields struct {
	Title    string   `url:"title"`
	Type     string   `url:"type"`
	SpaceId  string   `url:"spaceId,omitempty"`
	Shareds  []string `url:"shareds,omitempty,slice"`
	Settings Settings `url:"settings,omitempty,struct"`
}

type ModifyCustomFields struct {
	Title         string   `url:"title,omitempty"`
	Type          string   `url:"type,omitempty"`
	ChangeScope   string   `url:"changeScope,omitempty"`
	SpaceId       string   `url:"spaceId,omitempty"`
	AddShareds    []string `url:"addShareds,omitempty,slice"`
	RemoveShareds []string `url:"removeShareds,omitempty,slice"`
	Settings      Settings `url:"settings,omitempty,struct"`
}
