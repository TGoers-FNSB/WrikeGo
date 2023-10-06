package wrikeparams

type QueryCustomItemTypes struct {
	Title *string `url:"title,omitempty"`
	Limit *int    `url:"limit,omitempty"`
	Type  *string `url:"type,omitempty"`
}

type CreateWorkFromCustomItemType struct {
	SuperTaskId *string `url:"superTaskId,omitempty"`
	ParentId    *string `url:"parentId,omitempty"`
	Title       string  `url:"title"`
}
