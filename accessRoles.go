package wrikego

import (
	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryAccessRoles(config Config) (resp.AccessRoles, error) {
	path := "/access_roles"
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.AccessRolesFromJSON(response)
}