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
		Id            *string `json:"id,omitempty"`
		PredecessorId *string `json:"predecessorId,omitempty"`
		SuccessorId   *string `json:"successorId,omitempty"`
		RelationType  *string `json:"relationType,omitempty"`
	} `json:"data"`
}