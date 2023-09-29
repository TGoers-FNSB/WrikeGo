package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryApprovals(config Config, params params.QueryApprovals) (resp.Approvals, error) {
	path := "/approvals"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.ApprovalsFromJSON(response)
}

func QueryApprovalsByFolder(config Config, pathId string) (resp.Approvals, error) {
	path := fmt.Sprintf("/folders/%s/approvals", pathId)
	response, _ := Get(config, path, nil)
	return resp.ApprovalsFromJSON(response)
}

func QueryApprovalsByTask(config Config, pathId string) (resp.Approvals, error) {
	path := fmt.Sprintf("/tasks/%s/approvals", pathId)
	response, _ := Get(config, path, nil)
	return resp.ApprovalsFromJSON(response)
}

func QueryApprovalsByIds(config Config, pathId []string) (resp.Approvals, error) {
	path := fmt.Sprintf("/folders/%s/approvals", strings.Join(pathId, ","))
	response, _ := Get(config, path, nil)
	return resp.ApprovalsFromJSON(response)
}

func CreateApprovalsByFolder(config Config, params params.CreateApprovals, pathId string) (resp.Approvals, error) {
	path := fmt.Sprintf("/folders/%s/approvals", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.ApprovalsFromJSON(response)
}

func CreateApprovalsByTask(config Config, params params.CreateApprovals, pathId string) (resp.Approvals, error) {
	path := fmt.Sprintf("/tasks/%s/approvals", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.ApprovalsFromJSON(response)
}

func ModifyApprovalsById(config Config, params params.ModifyApprovals, pathId string) (resp.Approvals, error) {
	path := fmt.Sprintf("/approvals/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.ApprovalsFromJSON(response)
}

func DeleteApprovalsById(config Config, pathId string) (resp.Approvals, error) {
	path := fmt.Sprintf("/approvals/%s", pathId)
	response, _ := Delete(config, path, nil)
	return resp.ApprovalsFromJSON(response)
}