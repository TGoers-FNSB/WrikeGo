package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	types "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryTasks(config Config, params params.QueryTasks) (types.Tasks, error) {
	path := "/tasks"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return types.TasksFromJSON(response)
}

func QueryTasksInFolder(config Config, params params.QueryTasks, pathId string) (types.Tasks, error) {
	path := fmt.Sprintf("/folders/%s/tasks", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return types.TasksFromJSON(response)
}

func QueryTasksInSpace(config Config, params params.QueryTasks, pathId string) (types.Tasks, error) {
	path := fmt.Sprintf("/spaces/%s/tasks", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return types.TasksFromJSON(response)
}

func QueryTasksByIds(config Config, params params.QueryTasks, pathId []string) (types.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return types.TasksFromJSON(response)
}

func QueryTasksFieldsHistoryByIds(config Config, params params.QueryTasks, pathId []string) (types.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s/tasks_history", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return types.TasksFromJSON(response)
}

func CreateTaskInFolder(config Config, params params.CreateTasks, pathId string) (types.Tasks, error) {
	path := fmt.Sprintf("/folders/%s/tasks", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return types.TasksFromJSON(response)
}

func ModifyTaskById(config Config, params params.ModifyTasks, pathId string) (types.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return types.TasksFromJSON(response)
}

func ModifyTasksByIds(config Config, params params.ModifyTasks, pathId []string) (types.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return types.TasksFromJSON(response)
}

func DeleteTaskById(config Config, pathId string) (types.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", pathId)
	response, _ := Delete(config, path)
	return types.TasksFromJSON(response)
}