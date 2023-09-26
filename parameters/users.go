package wrikeparams

type ModifyUsers struct {
	Profile struct {
		AccountId string `url:"accountId"`
		Role      string `url:"profile"`
		External  *bool  `url:"external,omitempty"`
	} `url:"profile"`
}
