package wrikego

import (
	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryAccessRoles(config Config) (resp.AccessRoles, error) {
	path := "/access_roles"
	response, _ := Get(config, path, nil)
	return resp.AccessRolesFromJSON(response)
}