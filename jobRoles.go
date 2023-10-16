package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryJobRoles(config Config) (resp.JobRoles, *http.Response) {
	path := "/jobroles"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.JobRolesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryJobRolesByIds(config Config, pathId []string) (resp.JobRoles, *http.Response) {
	path := fmt.Sprintf("/jobroles/%s", strings.Join(pathId, ","))
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.JobRolesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateJobRoles(config Config, params params.CreateJobRoles) (resp.JobRoles, *http.Response) {
	path := "/jobroles"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.JobRolesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyJobRolesById(config Config, params params.ModifyJobRoles, pathId string) (resp.JobRoles, *http.Response) {
	path := fmt.Sprintf("/jobroles/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.JobRolesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteJobRolesById(config Config, pathId string) (resp.JobRoles, *http.Response) {
	path := fmt.Sprintf("/jobroles/%s", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.JobRolesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}