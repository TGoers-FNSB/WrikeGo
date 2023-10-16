package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryCustomFields(config Config) (resp.CustomFields, *http.Response) {
	path := "/customfields"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.CustomFieldsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryCustomFieldsByIds(config Config, pathId []string) (resp.CustomFields, *http.Response) {
	path := fmt.Sprintf("/customfields/%s", strings.Join(pathId, ","))
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.CustomFieldsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateCustomFields(config Config, params params.CreateCustomFields) (resp.CustomFields, *http.Response) {
	path := "/customfields"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.CustomFieldsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyCustomFieldsById(config Config, params params.ModifyCustomFields, pathId string) (resp.CustomFields, *http.Response) {
	path := fmt.Sprintf("/customfields/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.CustomFieldsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}