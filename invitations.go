package wrikego

import (
	"fmt"
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryInvitations(config Config) (resp.Invitations, *http.Response) {
	path := "/invitations"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.InvitationsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateInvitations(config Config, params params.CreateInvitations) (resp.Invitations, *http.Response) {
	path := "/invitations"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.InvitationsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyInvitationsById(config Config, params params.ModifyInvitaions, pathId string) (resp.Invitations, *http.Response) {
	path := fmt.Sprintf("/inivations/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.InvitationsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteInvitationsById(config Config, pathId string) (resp.Invitations, *http.Response) {
	path := fmt.Sprintf("/invitations/%s", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.InvitationsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}