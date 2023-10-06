package wrikeparams

type CreateJobRoles struct {
	Title       string  `url:"title"`
	ShortTitle  *string `url:"shortTitle,omitempty"`
	AvatarColor *string `url:"avatarColor,omitempty"`
}

type ModifyJobRoles struct {
	Title       *string `url:"title,omitempty"`
	ShortTitle  *string `url:"shortTitle,omitempty"`
	AvatarColor *string `url:"avatarColor,omitempty"`
}
