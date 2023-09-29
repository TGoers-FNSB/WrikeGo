package wrikeparams

type CreateEDiscovery struct {
	Scopes       []string  `url:"scopes"`
	Terms        *[]string `url:"terms,omitempty"`
	TargetUserId *string   `url:"targetUserId,omitempty"`
	Timeout      *string   `url:"timeout,omitempty"`
}
