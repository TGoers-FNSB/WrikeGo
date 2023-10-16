package wrikego

import (
	"fmt"
	
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryCustomItemTypes(config Config, params params.QueryCustomItemTypes) (resp.CustomItemTypes, error) {
	path := "/custom_item_types"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.CustomItemTypesFromJSON(response)
}

func QueryCustomItemTypesBySpace(config Config, params params.QueryCustomItemTypes, pathId string) (resp.CustomItemTypes, error) {
	path := fmt.Sprintf("/spaces/%s/custom_item_types", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.CustomItemTypesFromJSON(response)
}

func QueryCustomItemTypesByIds(config Config, pathId []string) (resp.CustomItemTypes, error) {
	path := fmt.Sprintf("/custom_item_types/%s", strings.Join(pathId, ","))
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.CustomItemTypesFromJSON(response)
}

func CreateWorkFromCustomItemTypesByCustomItemType(config Config, params params.CreateWorkFromCustomItemType, pathId string) (resp.CustomItemTypes, error) {
	path := fmt.Sprintf("/custom_item_types/%s/instantiate", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.CustomItemTypesFromJSON(response)
}