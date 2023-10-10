package wrikego

import (
	"fmt"
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryUsersById(config Config, pathId string) (resp.Users, error) {
	path := fmt.Sprintf("/users/%s", pathId)
	response, _ := Get(config, path, nil)
	return resp.UsersFromJSON(response)
}

func ModifyUsersById(config Config, params params.ModifyUsers, pathId string) (resp.Users, error) {
	path := fmt.Sprintf("/users/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.UsersFromJSON(response)
}