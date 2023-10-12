package wrikego

import (
	"fmt"
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QuerySpaces(config Config, params params.QuerySpaces) (resp.Spaces, error) {
	path := "/spaces"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.SpacesFromJSON(response)
}

func QuerySpacesById(config Config, params params.QuerySpaces, pathId string) (resp.Spaces, error) {
	path := fmt.Sprintf("/spaces/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.SpacesFromJSON(response)
}

func CreateSpaces(config Config, params params.CreateSpaces) (resp.Spaces, error) {
	path := "/spaces"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.SpacesFromJSON(response)
}

func ModifySpacesById(config Config, params params.ModifySpaces, pathId string) (resp.Spaces, error) {
	path := fmt.Sprintf("/spaces/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.SpacesFromJSON(response)
}

func DeleteSpacesById(config Config, pathId string) (resp.Spaces, error) {
	path := fmt.Sprintf("/spaces/%s", pathId)
	response, _ := Delete(config, path, nil)
	return resp.SpacesFromJSON(response)
}