package wrikego

import (
	"fmt"
	

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QuerySpaces(config Config, params params.QuerySpaces) (resp.Spaces, error) {
	path := "/spaces"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.SpacesFromJSON(response)
}

func QuerySpacesById(config Config, params params.QuerySpaces, pathId string) (resp.Spaces, error) {
	path := fmt.Sprintf("/spaces/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.SpacesFromJSON(response)
}

func CreateSpaces(config Config, params params.CreateSpaces) (resp.Spaces, error) {
	path := "/spaces"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.SpacesFromJSON(response)
}

func ModifySpacesById(config Config, params params.ModifySpaces, pathId string) (resp.Spaces, error) {
	path := fmt.Sprintf("/spaces/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.SpacesFromJSON(response)
}

func DeleteSpacesById(config Config, pathId string) (resp.Spaces, error) {
	path := fmt.Sprintf("/spaces/%s", pathId)
	response, err := Delete(config, path, nil)
	ErrorCheck(err)
	return resp.SpacesFromJSON(response)
}