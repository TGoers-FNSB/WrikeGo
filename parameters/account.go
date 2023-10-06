package wrikeparams

type QueryAccount struct {
	Metadata *Metadata `url:"metadata,omitempty"`
	Fields   *[]string `url:"fields,omitempty"`
}

type ModifyAccount struct {
	Metadata *[]Metadata `url:"metadata,omitempty"`
}
