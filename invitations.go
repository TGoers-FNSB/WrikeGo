package wrikego

import (
	"fmt"
	

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryInvitations(config Config) (resp.Invitations, error) {
	path := "/invitations"
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.InvitationsFromJSON(response)
}

func CreateInvitations(config Config, params params.CreateInvitations) (resp.Invitations, error) {
	path := "/invitations"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.InvitationsFromJSON(response)
}

func ModifyInvitationsById(config Config, params params.ModifyInvitaions, pathId string) (resp.Invitations, error) {
	path := fmt.Sprintf("/inivations/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.InvitationsFromJSON(response)
}

func DeleteInvitationsById(config Config, pathId string) (resp.Invitations, error) {
	path := fmt.Sprintf("/invitations/%s", pathId)
	response, err := Delete(config, path, nil)
	ErrorCheck(err)
	return resp.InvitationsFromJSON(response)
}