package wrikego

import (
	"fmt"
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryGroupsById(config Config, params params.QueryGroups, pathId string) (resp.Groups, error) {
	path := fmt.Sprintf("/groups/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.GroupsFromJSON(response)
}

func QueryGroups(config Config, params params.QueryGroups) (resp.Groups, error) {
	path := "/groups"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.GroupsFromJSON(response)
}

func CreateGroups(config Config, params params.CreateGroups) (resp.Groups, error) {
	path := "/groups"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.GroupsFromJSON(response)
}

func ModifyGroupsBulk(config Config, params params.ModifyGroupsBulk) (resp.Groups, error) {
	path := "groups_bulk"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.GroupsFromJSON(response)
}

func ModifyGroupsById(config Config, params params.ModifyGroups, pathId string) (resp.Groups, error) {
	path := fmt.Sprintf("/groups/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.GroupsFromJSON(response)
}

func DeleteGroupsById(config Config, params params.DeleteGroups, pathId string) (resp.Groups, error) {
	path := fmt.Sprintf("/groups/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Delete(config, path, body)
	return resp.GroupsFromJSON(response)
}