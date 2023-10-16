package wrikego

import (
	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryUserTypes(config Config) (resp.UserTypes, error) {
	path := "/user_types"
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.UserTypesFromJSON(response)
}
