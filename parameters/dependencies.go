package wrikeparams

type CreateDependencies struct {
	PredecessorId string `url:"predecessorId"`
	SuccessorId   string `url:"successorId"`
	RelationType  string `url:"relationType"`
}

type ModifyDependencies struct {
	RelationType string `url:"relationType"`
}
