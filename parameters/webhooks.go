package wrikeparams

type CreateWebooks struct {
	HookUrl   string   `url:"hookUrl"`
	Secret    string   `url:"secret,omitempty"`
	Events    []string `url:"events,omitempty,slice"`
	Recursive bool     `url:"recursive,omitempty"`
}

type ModifyWebhooks struct {
	Status string `url:"status"`
}
