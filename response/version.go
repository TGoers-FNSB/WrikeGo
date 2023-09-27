package wrikeresponse

import "encoding/json"

func VersionFromJSON(data []byte) (Version, error) {
	var item Version
	err := json.Unmarshal(data, &item)
	return item, err
}

type Version struct {
	Kind string `json:"kind"`
	Data []struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
	} `json:"data"`
}