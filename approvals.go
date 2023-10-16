package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryApprovals(config Config, params params.QueryApprovals) (resp.Approvals, *http.Response) {
	path := "/approvals"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.ApprovalsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryApprovalsByFolder(config Config, pathId string) (resp.Approvals, *http.Response) {
	path := fmt.Sprintf("/folders/%s/approvals", pathId)
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.ApprovalsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryApprovalsByTask(config Config, pathId string) (resp.Approvals, *http.Response) {
	path := fmt.Sprintf("/tasks/%s/approvals", pathId)
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.ApprovalsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryApprovalsByIds(config Config, pathId []string) (resp.Approvals, *http.Response) {
	path := fmt.Sprintf("/folders/%s/approvals", strings.Join(pathId, ","))
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.ApprovalsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateApprovalsByFolder(config Config, params params.CreateApprovals, pathId string) (resp.Approvals, *http.Response) {
	path := fmt.Sprintf("/folders/%s/approvals", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.ApprovalsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateApprovalsByTask(config Config, params params.CreateApprovals, pathId string) (resp.Approvals, *http.Response) {
	path := fmt.Sprintf("/tasks/%s/approvals", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.ApprovalsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyApprovalsById(config Config, params params.ModifyApprovals, pathId string) (resp.Approvals, *http.Response) {
	path := fmt.Sprintf("/approvals/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.ApprovalsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteApprovalsById(config Config, pathId string) (resp.Approvals, *http.Response) {
	path := fmt.Sprintf("/approvals/%s", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.ApprovalsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}