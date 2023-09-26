package wrikeparams

type CreateInvitations struct {
	Email     string  `url:"email"`
	FirstName *string `url:"firstName,omitempty"`
	LastName  *string `url:"lastName,omitempty"`
	Role      *string `url:"role,omitempty"`
	External  *bool   `url:"external,omitempty"`
	Subject   *string `url:"subject,omitempty"`
	Message   *string `url:"message,omitempty"`
}

type ModifyInvitaions struct {
	Resend   *string `url:"resend,omitempty"`
	Role     *string `url:"role,omitempty"`
	External *bool   `url:"external,omitempty"`
}
