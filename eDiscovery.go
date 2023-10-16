package wrikego

import (
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func CreateEDiscovery(config Config, params params.CreateEDiscovery, pathId string) (resp.EDiscovery, *http.Response) {
	path := "/ediscovery_search"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.EDiscoveryFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}