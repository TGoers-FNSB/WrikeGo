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
		Id             *string `json:"id,omitempty"`
		SpaceId        *string `json:"spaceId,omitempty"`
		Name           *string `json:"name,omitempty"`
		Standard       *bool   `json:"standard,omitempty"`
		Hidden         *bool   `json:"hidden,omitempty"`
		CustomStatuses *[]struct {
			Id           *string `json:"id,omitempty"`
			Name         *string `json:"name,omitempty"`
			StandardName *bool   `json:"standardName,omitempty"`
			Color        *string `json:"color,omitempty"`
			Standard     *bool   `json:"standard,omitempty"`
			Group        *string `json:"group,omitempty"`
			Hidden       *bool   `json:"hidden,omitempty"`
		} `json:"customStatuses,omitempty"`
	} `json:"data"`
}
