package wrikego

import (
	"fmt"
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryWorkflows(config Config) (resp.Workflows, error) {
	path := "/workflows"
	response, _ := Get(config, path, nil)
	return resp.WorkflowsFromJSON(response)
}

func CreateWorkflows(config Config, params params.CreateWorkflows) (resp.Workflows, error) {
	path := "/workflows"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.WorkflowsFromJSON(response)
}

func ModifyWorkflowsById(config Config, params params.CreateWorkflows, pathId string) (resp.Workflows, error) {
	path := fmt.Sprintf("/workflows/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.WorkflowsFromJSON(response)
}