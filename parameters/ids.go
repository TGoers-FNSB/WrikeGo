package wrikeparams

type QueryIDs struct {
	Type string   `url:"type"`
	Ids  []string `url:"ids"`
}
