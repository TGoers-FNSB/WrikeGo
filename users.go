package wrikego

import (
	"fmt"
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryUsersById(config Config, pathId string) (resp.Users, *http.Response) {
	path := fmt.Sprintf("/users/%s", pathId)
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.UsersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyUsersById(config Config, params params.ModifyUsers, pathId string) (resp.Users, *http.Response) {
	path := fmt.Sprintf("/users/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.UsersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}