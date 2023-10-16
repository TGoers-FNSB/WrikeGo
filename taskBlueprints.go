package wrikego

import (
	"fmt"
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryTaskBlueprints(config Config, params params.QueryTaskBlueprints) (resp.TaskBlueprints, *http.Response) {
	path := "/task_blueprints"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.TaskBlueprintsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryTaskBlueprintsBySpace(config Config, params params.QueryTaskBlueprints, pathId string) (resp.TaskBlueprints, *http.Response) {
	path := fmt.Sprintf("/spaces/%s/task_blueprints", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.TaskBlueprintsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateTaskBlueprintsAsync(config Config, params params.CreateTaskBlueprintsAsync, pathId string) (resp.TaskBlueprints, *http.Response) {
	path := fmt.Sprintf("/task_blueprints/%s/launch_async", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.TaskBlueprintsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}