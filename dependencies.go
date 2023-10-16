package wrikego

import (
	"fmt"
	
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryDependenciesByTask(config Config, pathId string) (resp.Dependencies, error) {
	path := fmt.Sprintf("/tasks/%s/dependencies", pathId)
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.DependenciesFromJSON(response)
}

func QueryDependenciesByIds(config Config, pathId []string) (resp.Dependencies, error) {
	path := fmt.Sprintf("/dependencies/%s", strings.Join(pathId, ","))
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.DependenciesFromJSON(response)
}

func CreateDependenciesByTask(config Config, params params.CreateDependencies, pathId string) (resp.Dependencies, error) {
	path := fmt.Sprintf("/tasks/%s/dependencies", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.DependenciesFromJSON(response)
}

func ModifyDependenciesById(config Config, params params.ModifyDependencies, pathId string) (resp.Dependencies, error) {
	path := fmt.Sprintf("/dependencies/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.DependenciesFromJSON(response)
}

func DeleteDependenciesById(config Config, pathId string) (resp.Dependencies, error) {
	path := fmt.Sprintf("/dependencies/%s", pathId)
	response, err := Delete(config, path, nil)
	ErrorCheck(err)
	return resp.DependenciesFromJSON(response)
}