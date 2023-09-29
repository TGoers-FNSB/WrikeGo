package wrikego

import (
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func CreateEDiscovery(config Config, params params.CreateEDiscovery, pathId string) (resp.EDiscovery, error) {
	path := "/ediscovery_search"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.EDiscoveryFromJSON(response)
}