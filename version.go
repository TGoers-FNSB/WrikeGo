package wrikego

import (
	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryVersion(config Config) (resp.Version, error) {
	path := "/version"
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.VersionFromJSON(response)
}