package wrikego

import (
	"fmt"
	
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryJobRoles(config Config) (resp.JobRoles, error) {
	path := "/jobroles"
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.JobRolesFromJSON(response)
}

func QueryJobRolesByIds(config Config, pathId []string) (resp.JobRoles, error) {
	path := fmt.Sprintf("/jobroles/%s", strings.Join(pathId, ","))
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.JobRolesFromJSON(response)
}

func CreateJobRoles(config Config, params params.CreateJobRoles) (resp.JobRoles, error) {
	path := "/jobroles"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.JobRolesFromJSON(response)
}

func ModifyJobRolesById(config Config, params params.ModifyJobRoles, pathId string) (resp.JobRoles, error) {
	path := fmt.Sprintf("/jobroles/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.JobRolesFromJSON(response)
}

func DeleteJobRolesById(config Config, pathId string) (resp.JobRoles, error) {
	path := fmt.Sprintf("/jobroles/%s", pathId)
	response, err := Delete(config, path, nil)
	ErrorCheck(err)
	return resp.JobRolesFromJSON(response)
}