package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryTasks(config Config, params params.QueryTasks) (resp.Tasks, *http.Response) {
	path := "/tasks"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.TasksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryTasksByFolder(config Config, params params.QueryTasks, pathId string) (resp.Tasks, *http.Response) {
	path := fmt.Sprintf("/folders/%s/tasks", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.TasksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryTasksBySpace(config Config, params params.QueryTasks, pathId string) (resp.Tasks, *http.Response) {
	path := fmt.Sprintf("/spaces/%s/tasks", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.TasksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryTasksByIds(config Config, params params.QueryTasks, pathId []string) (resp.Tasks, *http.Response) {
	path := fmt.Sprintf("/tasks/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.TasksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryTasksFieldsHistoryByTasks(config Config, params params.QueryTasksFieldsHistory, pathId []string) (resp.Tasks, *http.Response) {
	path := fmt.Sprintf("/tasks/%s/tasks_history", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.TasksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}
//----------------------
func CreateTasksByFolder(config Config, params params.CreateTasks, pathId string) (resp.Tasks, *http.Response) {
	path := fmt.Sprintf("/folders/%s/tasks", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.TasksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}
//-----------------------
func ModifyTasksById(config Config, params params.ModifyTasks, pathId string) (resp.Tasks, *http.Response) {
	path := fmt.Sprintf("/tasks/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.TasksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyTasksByIds(config Config, params params.ModifyTasks, pathId []string) (resp.Tasks, *http.Response) {
	path := fmt.Sprintf("/tasks/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.TasksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteTasksById(config Config, pathId string) (resp.Tasks, *http.Response) {
	path := fmt.Sprintf("/tasks/%s", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.TasksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}