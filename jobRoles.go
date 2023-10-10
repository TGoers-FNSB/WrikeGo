package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryJobRoles(config Config) (resp.JobRoles, error) {
	path := "/jobroles"
	response, _ := Get(config, path, nil)
	return resp.JobRolesFromJSON(response)
}

func QueryJobRolesByIds(config Config, pathId []string) (resp.JobRoles, error) {
	path := fmt.Sprintf("/jobroles/%s", strings.Join(pathId, ","))
	response, _ := Get(config, path, nil)
	return resp.JobRolesFromJSON(response)
}

func CreateJobRoles(config Config, params params.CreateJobRoles) (resp.JobRoles, error) {
	path := "/jobroles"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.JobRolesFromJSON(response)
}

func ModifyJobRolesById(config Config, params params.ModifyJobRoles, pathId string) (resp.JobRoles, error) {
	path := fmt.Sprintf("/jobroles/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.JobRolesFromJSON(response)
}

func DeleteJobRolesById(config Config, pathId string) (resp.JobRoles, error) {
	path := fmt.Sprintf("/jobroles/%s", pathId)
	response, _ := Delete(config, path, nil)
	return resp.JobRolesFromJSON(response)
}