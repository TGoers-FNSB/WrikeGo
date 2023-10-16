package wrikego

import (
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryAccount(config Config, params params.QueryAccount) (resp.Account, *http.Response) {
	path := "/account"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.AccountFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyAccount(config Config, params params.QueryAccount) (resp.Account, *http.Response) {
	path := "/account"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.AccountFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}