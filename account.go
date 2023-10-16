package wrikego

import (
	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryAccount(config Config, params params.QueryAccount) (resp.Account, error) {
	path := "/account"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.AccountFromJSON(response)
}

func ModifyAccount(config Config, params params.QueryAccount) (resp.Account, error) {
	path := "/account"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.AccountFromJSON(response)
}