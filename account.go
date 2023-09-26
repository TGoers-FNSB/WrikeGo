package wrikego

import (
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryAccount(config Config, params params.QueryAccount) (resp.Account, error) {
	path := "/account"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.AccountFromJSON(response)
}

func ModifyAccount(config Config, params params.QueryAccount) (resp.Account, error) {
	path := "/account"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.AccountFromJSON(response)
}