package wrikeparams

type CreateEDiscovery struct {
	Scopes       []string `url:"scopes,slice"`
	Terms        []string `url:"terms,omitempty,slice"`
	TargetUserId string   `url:"targetUserId,omitempty"`
	Timeout      string   `url:"timeout,omitempty"`
}
