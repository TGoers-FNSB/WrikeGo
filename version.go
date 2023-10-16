package wrikego

import (
	"net/http"

	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryVersion(config Config) (resp.Version, *http.Response) {
	path := "/version"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.VersionFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}