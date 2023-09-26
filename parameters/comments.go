package wrikeparams

type QueryComments struct {
	PlainText   *bool        `url:"plainText,omitempty"`
	Types       *[]string    `url:"types,omitempty"`
	UpdatedDate *DateOrRange `url:"updatedDate,omitempty"`
	Limit       *int         `url:"limit,omitempty"`
	Fields      *[]string    `url:"fields,omitempty"`
}

type CreateComments struct {
	Text      string `url:"text"`
	PlainText *bool  `url:"plainText,omitempty"`
}

type ModifyComments struct {
	Text      string `url:"text"`
	PlainText *bool  `url:"plainText,omitempty"`
}
