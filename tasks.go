package wrikego

import (
	"fmt"
	
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryTasks(config Config, params params.QueryTasks) (resp.Tasks, error) {
	path := "/tasks"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.TasksFromJSON(response)
}

func QueryTasksByFolder(config Config, params params.QueryTasks, pathId string) (resp.Tasks, error) {
	path := fmt.Sprintf("/folders/%s/tasks", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.TasksFromJSON(response)
}

func QueryTasksBySpace(config Config, params params.QueryTasks, pathId string) (resp.Tasks, error) {
	path := fmt.Sprintf("/spaces/%s/tasks", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.TasksFromJSON(response)
}

func QueryTasksByIds(config Config, params params.QueryTasks, pathId []string) (resp.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.TasksFromJSON(response)
}

func QueryTasksFieldsHistoryByTasks(config Config, params params.QueryTasksFieldsHistory, pathId []string) (resp.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s/tasks_history", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.TasksFromJSON(response)
}
//----------------------
func CreateTasksByFolder(config Config, params params.CreateTasks, pathId string) (resp.Tasks, error) {
	path := fmt.Sprintf("/folders/%s/tasks", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.TasksFromJSON(response)
}
//-----------------------
func ModifyTasksById(config Config, params params.ModifyTasks, pathId string) (resp.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.TasksFromJSON(response)
}

func ModifyTasksByIds(config Config, params params.ModifyTasks, pathId []string) (resp.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.TasksFromJSON(response)
}

func DeleteTasksById(config Config, pathId string) (resp.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", pathId)
	response, err := Delete(config, path, nil)
	ErrorCheck(err)
	return resp.TasksFromJSON(response)
}