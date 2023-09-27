package wrikego

import (
	"fmt"
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
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

func QuerySpaceById(config Config, params params.QuerySpaces, pathId string) (resp.Spaces, error) {
	path := fmt.Sprintf("/spaces/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.SpacesFromJSON(response)
}

func CreateSpace(config Config, params params.CreateSpaces) (resp.Spaces, error) {
	path := "/spaces"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.SpacesFromJSON(response)
}

func ModifySpaceById(config Config, params params.ModifySpaces, pathId string) (resp.Spaces, error) {
	path := fmt.Sprintf("/spaces/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.SpacesFromJSON(response)
}

func DeleteSpaceById(config Config, pathId string) (resp.Spaces, error) {
	path := fmt.Sprintf("/spaces/%s", pathId)
	response, _ := Delete(config, path, nil)
	return resp.SpacesFromJSON(response)
}