package wrikego

import (
	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryVersion(config Config) (resp.Version, error) {
	path := "/version"
	response, _ := Get(config, path, nil)
	return resp.VersionFromJSON(response)
}