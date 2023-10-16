package wrikego

import (
	"fmt"
	

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryUsersById(config Config, pathId string) (resp.Users, error) {
	path := fmt.Sprintf("/users/%s", pathId)
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.UsersFromJSON(response)
}

func ModifyUsersById(config Config, params params.ModifyUsers, pathId string) (resp.Users, error) {
	path := fmt.Sprintf("/users/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.UsersFromJSON(response)
}