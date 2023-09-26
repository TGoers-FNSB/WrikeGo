package wrikeresponse

import "encoding/json"

func WorkflowsFromJSON(data []byte) (Workflows, error) {
	var item Workflows
	err := json.Unmarshal(data, &item)
	return item, err
}

type Workflows struct {
	Kind string `json:"kind"`
	Data []struct {
		Id             string `json:"id"`
		SpaceId			string	`json:"spaceId"`
		Name           string `json:"name"`
		Standard       bool   `json:"standard"`
		Hidden         bool   `json:"hidden"`
		CustomStatuses []struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			StandardName bool   `json:"standardName"`
			Color        string `json:"color"`
			Standard     bool   `json:"standard"`
			Group        string `json:"group"`
			Hidden       bool   `json:"hidden"`
		} `json:"customStatuses"`
	} `json:"data"`
}