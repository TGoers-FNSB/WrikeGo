package wrikego

import (
	"fmt"
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryInvitations(config Config) (resp.Invitations, error) {
	path := "/invitations"
	response, _ := Get(config, path, nil)
	return resp.InvitationsFromJSON(response)
}

func CreateInvitations(config Config, params params.CreateInvitations) (resp.Invitations, error) {
	path := "/invitations"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.InvitationsFromJSON(response)
}

func ModifyInvitationsById(config Config, params params.ModifyInvitaions, pathId string) (resp.Invitations, error) {
	path := fmt.Sprintf("/inivations/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.InvitationsFromJSON(response)
}

func DeleteInvitationsById(config Config, pathId string) (resp.Invitations, error) {
	path := fmt.Sprintf("/invitations/%s", pathId)
	response, _ := Delete(config, path, nil)
	return resp.InvitationsFromJSON(response)
}