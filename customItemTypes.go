package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryCustomItemTypes(config Config, params params.QueryCustomItemTypes) (resp.CustomItemTypes, *http.Response) {
	path := "/custom_item_types"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.CustomItemTypesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryCustomItemTypesBySpace(config Config, params params.QueryCustomItemTypes, pathId string) (resp.CustomItemTypes, *http.Response) {
	path := fmt.Sprintf("/spaces/%s/custom_item_types", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.CustomItemTypesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryCustomItemTypesByIds(config Config, pathId []string) (resp.CustomItemTypes, *http.Response) {
	path := fmt.Sprintf("/custom_item_types/%s", strings.Join(pathId, ","))
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.CustomItemTypesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateWorkFromCustomItemTypesByCustomItemType(config Config, params params.CreateWorkFromCustomItemType, pathId string) (resp.CustomItemTypes, *http.Response) {
	path := fmt.Sprintf("/custom_item_types/%s/instantiate", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.CustomItemTypesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}