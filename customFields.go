package wrikego

import (
	"fmt"
	
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryCustomFields(config Config) (resp.CustomFields, error) {
	path := "/customfields"
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.CustomFieldsFromJSON(response)
}

func QueryCustomFieldsByIds(config Config, pathId []string) (resp.CustomFields, error) {
	path := fmt.Sprintf("/customfields/%s", strings.Join(pathId, ","))
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.CustomFieldsFromJSON(response)
}

func CreateCustomFields(config Config, params params.CreateCustomFields) (resp.CustomFields, error) {
	path := "/customfields"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.CustomFieldsFromJSON(response)
}

func ModifyCustomFieldsById(config Config, params params.ModifyCustomFields, pathId string) (resp.CustomFields, error) {
	path := fmt.Sprintf("/customfields/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.CustomFieldsFromJSON(response)
}