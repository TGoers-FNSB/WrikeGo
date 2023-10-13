package wrikeparams

type QueryAccount struct {
	Metadata Metadata `url:"metadata,omitempty,struct"`
	Fields   []string `url:"fields,omitempty,slice"`
}

type ModifyAccount struct {
	Metadata []Metadata `url:"metadata,omitempty,slice+struct"`
}
