package wrikeparams

type CreateJobRoles struct {
	Title       string  `url:"title"`
	ShortTitle  *string `url:"shortTitle"`
	AvatarColor *string `url:"avatarColor"`
}

type ModifyJobRoles struct {
	Title       *string `url:"title"`
	ShortTitle  *string `url:"shortTitle"`
	AvatarColor *string `url:"avatarColor"`
}
