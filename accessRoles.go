package wrikego

import (
	"net/http"

	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryAccessRoles(config Config) (resp.AccessRoles, *http.Response) {
	path := "/access_roles"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.AccessRolesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}