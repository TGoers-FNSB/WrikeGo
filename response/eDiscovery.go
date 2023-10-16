package wrikeresponse

import (
	"encoding/json"
)

func EDiscoveryFromJSON(data []byte) (EDiscovery, error) {
	var item EDiscovery
	err := json.Unmarshal(data, &item)
	return item, err
}

type EDiscovery struct {
	Kind string `json:"kind"`
	Data []struct {
		Type *string `json:"type,omitempty"`
		Id   *string `json:"id,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}