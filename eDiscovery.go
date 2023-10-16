package wrikego

import (
	

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func CreateEDiscovery(config Config, params params.CreateEDiscovery, pathId string) (resp.EDiscovery, error) {
	path := "/ediscovery_search"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.EDiscoveryFromJSON(response)
}