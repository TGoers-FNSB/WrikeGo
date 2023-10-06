package wrikego

import (
	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryUserTypes(config Config) (resp.UserTypes, error) {
	path := "/user_types"
	response, _ := Get(config, path, nil)
	return resp.UserTypesFromJSON(response)
}
