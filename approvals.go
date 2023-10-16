package wrikego

import (
	"fmt"
	
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryApprovals(config Config, params params.QueryApprovals) (resp.Approvals, error) {
	path := "/approvals"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.ApprovalsFromJSON(response)
}

func QueryApprovalsByFolder(config Config, pathId string) (resp.Approvals, error) {
	path := fmt.Sprintf("/folders/%s/approvals", pathId)
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.ApprovalsFromJSON(response)
}

func QueryApprovalsByTask(config Config, pathId string) (resp.Approvals, error) {
	path := fmt.Sprintf("/tasks/%s/approvals", pathId)
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.ApprovalsFromJSON(response)
}

func QueryApprovalsByIds(config Config, pathId []string) (resp.Approvals, error) {
	path := fmt.Sprintf("/folders/%s/approvals", strings.Join(pathId, ","))
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.ApprovalsFromJSON(response)
}

func CreateApprovalsByFolder(config Config, params params.CreateApprovals, pathId string) (resp.Approvals, error) {
	path := fmt.Sprintf("/folders/%s/approvals", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.ApprovalsFromJSON(response)
}

func CreateApprovalsByTask(config Config, params params.CreateApprovals, pathId string) (resp.Approvals, error) {
	path := fmt.Sprintf("/tasks/%s/approvals", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.ApprovalsFromJSON(response)
}

func ModifyApprovalsById(config Config, params params.ModifyApprovals, pathId string) (resp.Approvals, error) {
	path := fmt.Sprintf("/approvals/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.ApprovalsFromJSON(response)
}

func DeleteApprovalsById(config Config, pathId string) (resp.Approvals, error) {
	path := fmt.Sprintf("/approvals/%s", pathId)
	response, err := Delete(config, path, nil)
	ErrorCheck(err)
	return resp.ApprovalsFromJSON(response)
}