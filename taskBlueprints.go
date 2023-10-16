package wrikego

import (
	"fmt"
	

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryTaskBlueprints(config Config, params params.QueryTaskBlueprints) (resp.TaskBlueprints, error) {
	path := "/task_blueprints"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.TaskBlueprintsFromJSON(response)
}

func QueryTaskBlueprintsBySpace(config Config, params params.QueryTaskBlueprints, pathId string) (resp.TaskBlueprints, error) {
	path := fmt.Sprintf("/spaces/%s/task_blueprints", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.TaskBlueprintsFromJSON(response)
}

func CreateTaskBlueprintsAsync(config Config, params params.CreateTaskBlueprintsAsync, pathId string) (resp.TaskBlueprints, error) {
	path := fmt.Sprintf("/task_blueprints/%s/launch_async", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.TaskBlueprintsFromJSON(response)
}