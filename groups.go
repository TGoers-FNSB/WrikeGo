package wrikego

import (
	"fmt"
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryGroupsById(config Config, params params.QueryGroups, pathId string) (resp.Groups, *http.Response) {
	path := fmt.Sprintf("/groups/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.GroupsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryGroups(config Config, params params.QueryGroups) (resp.Groups, *http.Response) {
	path := "/groups"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.GroupsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateGroups(config Config, params params.CreateGroups) (resp.Groups, *http.Response) {
	path := "/groups"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.GroupsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyGroupsBulk(config Config, params params.ModifyGroupsBulk) (resp.Groups, *http.Response) {
	path := "groups_bulk"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.GroupsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyGroupsById(config Config, params params.ModifyGroups, pathId string) (resp.Groups, *http.Response) {
	path := fmt.Sprintf("/groups/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.GroupsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteGroupsById(config Config, params params.DeleteGroups, pathId string) (resp.Groups, *http.Response) {
	path := fmt.Sprintf("/groups/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Delete(config, path, body)
	ErrorCheck(err)
	json, err := resp.GroupsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}