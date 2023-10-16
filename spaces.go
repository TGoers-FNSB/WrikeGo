package wrikego

import (
	"fmt"
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QuerySpaces(config Config, params params.QuerySpaces) (resp.Spaces, *http.Response) {
	path := "/spaces"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.SpacesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QuerySpacesById(config Config, params params.QuerySpaces, pathId string) (resp.Spaces, *http.Response) {
	path := fmt.Sprintf("/spaces/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.SpacesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateSpaces(config Config, params params.CreateSpaces) (resp.Spaces, *http.Response) {
	path := "/spaces"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.SpacesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifySpacesById(config Config, params params.ModifySpaces, pathId string) (resp.Spaces, *http.Response) {
	path := fmt.Sprintf("/spaces/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.SpacesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteSpacesById(config Config, pathId string) (resp.Spaces, *http.Response) {
	path := fmt.Sprintf("/spaces/%s", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.SpacesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}