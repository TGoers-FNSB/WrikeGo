package wrikego

import (
	"fmt"
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryTaskBlueprints(config Config, params params.QueryTaskBlueprints) (resp.TaskBlueprints, error) {
	path := "/task_blueprints"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.TaskBlueprintsFromJSON(response)
}

func CreateTaskBlueprintsAsync(config Config, params params.CreateTaskBlueprintsAsync, pathId string) (resp.TaskBlueprints, error) {
	path := fmt.Sprintf("/task_blueprints/%s/launch_async", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.TaskBlueprintsFromJSON(response)
}