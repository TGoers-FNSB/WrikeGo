package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryDependenciesByTask(config Config, pathId string) (resp.Dependencies, *http.Response) {
	path := fmt.Sprintf("/tasks/%s/dependencies", pathId)
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.DependenciesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryDependenciesByIds(config Config, pathId []string) (resp.Dependencies, *http.Response) {
	path := fmt.Sprintf("/dependencies/%s", strings.Join(pathId, ","))
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.DependenciesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateDependenciesByTask(config Config, params params.CreateDependencies, pathId string) (resp.Dependencies, *http.Response) {
	path := fmt.Sprintf("/tasks/%s/dependencies", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.DependenciesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyDependenciesById(config Config, params params.ModifyDependencies, pathId string) (resp.Dependencies, *http.Response) {
	path := fmt.Sprintf("/dependencies/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.DependenciesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteDependenciesById(config Config, pathId string) (resp.Dependencies, *http.Response) {
	path := fmt.Sprintf("/dependencies/%s", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.DependenciesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}