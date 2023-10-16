package wrikego

import (
	"fmt"
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryWorkflows(config Config) (resp.Workflows, *http.Response) {
	path := "/workflows"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.WorkflowsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateWorkflows(config Config, params params.CreateWorkflows) (resp.Workflows, *http.Response) {
	path := "/workflows"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.WorkflowsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyWorkflowsById(config Config, params params.CreateWorkflows, pathId string) (resp.Workflows, *http.Response) {
	path := fmt.Sprintf("/workflows/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.WorkflowsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}