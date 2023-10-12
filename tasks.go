package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)


func Parse(jsonString string) (string) {
	test := ""

	fmt.Println("Test:", test)

	return ""
}

func QueryTasks(config Config, params params.QueryTasks) (resp.Tasks, error) {
	path := "/tasks"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.TasksFromJSON(response)
}

// Working out how to encode slices properly and objects properly

func QueryTasksByFolder(config Config, params params.QueryTasks, pathId string) (resp.Tasks, error) {
	path := fmt.Sprintf("/folders/%s/tasks", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.TasksFromJSON(response)
}

func QueryTasksBySpace(config Config, params params.QueryTasks, pathId string) (resp.Tasks, error) {
	path := fmt.Sprintf("/spaces/%s/tasks", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.TasksFromJSON(response)
}

func QueryTasksByIds(config Config, params params.QueryTasks, pathId []string) (resp.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.TasksFromJSON(response)
}

func QueryTasksFieldsHistoryByTasks(config Config, params params.QueryTasksFieldsHistory, pathId []string) (resp.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s/tasks_history", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.TasksFromJSON(response)
}
//----------------------
func CreateTasksByFolder(config Config, params params.CreateTasks, pathId string) (resp.Tasks, error) {
	path := fmt.Sprintf("/folders/%s/tasks", pathId)
	body, err := Values(params)




	
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.TasksFromJSON(response)




	// path := fmt.Sprintf("/folders/%s/tasks", pathId)
	// body, err := query.Values(params)
	// if err != nil {
	// 	log.Println(err)
	// }
	// response, _ := Post(config, path, body)
	// return resp.TasksFromJSON(response)
}
//-----------------------
func ModifyTasksById(config Config, params params.ModifyTasks, pathId string) (resp.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.TasksFromJSON(response)
}

func ModifyTasksByIds(config Config, params params.ModifyTasks, pathId []string) (resp.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.TasksFromJSON(response)
}

func DeleteTasksById(config Config, pathId string) (resp.Tasks, error) {
	path := fmt.Sprintf("/tasks/%s", pathId)
	response, _ := Delete(config, path, nil)
	return resp.TasksFromJSON(response)
}