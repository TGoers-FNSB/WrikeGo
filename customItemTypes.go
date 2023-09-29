package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryCustomItemTypes(config Config, params params.QueryCustomItemTypes) (resp.CustomItemTypes, error) {
	path := "/custom_item_types"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.CustomItemTypesFromJSON(response)
}

func QueryCustomItemTypesBySpace(config Config, params params.QueryCustomItemTypes, pathId string) (resp.CustomItemTypes, error) {
	path := fmt.Sprintf("/spaces/%s/custom_item_types", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.CustomItemTypesFromJSON(response)
}

func QueryCustomItemTypesByIds(config Config, pathId []string) (resp.CustomItemTypes, error) {
	path := fmt.Sprintf("/custom_item_types/%s", strings.Join(pathId, ","))
	response, _ := Get(config, path, nil)
	return resp.CustomItemTypesFromJSON(response)
}

func CreateWorkFromCustomItemTypeByCustomItemType(config Config, params params.CreateWorkFromCustomItemType, pathId string) (resp.CustomItemTypes, error) {
	path := fmt.Sprintf("/custom_item_types/%s/instantiate", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.CustomItemTypesFromJSON(response)
}