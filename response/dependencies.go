package wrikeresponse

import "encoding/json"

func DependenciesFromJSON(data []byte) (Dependencies, error) {
	var item Dependencies
	err := json.Unmarshal(data, &item)
	return item, err
}

type Dependencies struct {
	Kind string `json:"kind"`
	Data []struct {
		Id            string `json:"id"`
		PredecessorId string `json:"predecessorId"`
		SuccessorId   string `json:"successorId"`
		RelationType  string `json:"relationType"`
	} `json:"data"`
}