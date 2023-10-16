package wrikego

import (
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryIDs(config Config, params params.QueryIDs) (resp.IDs, *http.Response) {
	path := "/ids"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.IDsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}