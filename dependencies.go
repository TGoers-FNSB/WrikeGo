package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryDependencyByTask(config Config, pathId string) (resp.Dependencies, error) {
	path := fmt.Sprintf("/tasks/%s/dependencies", pathId)
	response, _ := Get(config, path, nil)
	return resp.DependenciesFromJSON(response)
}

func QueryDependenciesByIds(config Config, pathId []string) (resp.Dependencies, error) {
	path := fmt.Sprintf("/dependencies/%s", strings.Join(pathId, ","))
	response, _ := Get(config, path, nil)
	return resp.DependenciesFromJSON(response)
}

func CreateDependencyByTask(config Config, params params.CreateDependencies, pathId string) (resp.Dependencies, error) {
	path := fmt.Sprintf("/tasks/%s/dependencies", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.DependenciesFromJSON(response)
}

func ModifyDependencyById(config Config, params params.ModifyDependencies, pathId string) (resp.Dependencies, error) {
	path := fmt.Sprintf("/dependencies/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.DependenciesFromJSON(response)
}

func DeleteDependencyById(config Config, pathId string) (resp.Dependencies, error) {
	path := fmt.Sprintf("/dependencies/%s", pathId)
	response, _ := Delete(config, path, nil)
	return resp.DependenciesFromJSON(response)
}