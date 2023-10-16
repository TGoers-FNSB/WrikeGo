package wrikego

import (
	"fmt"
	

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryWorkflows(config Config) (resp.Workflows, error) {
	path := "/workflows"
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.WorkflowsFromJSON(response)
}

func CreateWorkflows(config Config, params params.CreateWorkflows) (resp.Workflows, error) {
	path := "/workflows"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.WorkflowsFromJSON(response)
}

func ModifyWorkflowsById(config Config, params params.CreateWorkflows, pathId string) (resp.Workflows, error) {
	path := fmt.Sprintf("/workflows/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.WorkflowsFromJSON(response)
}