package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryCustomFields(config Config) (resp.CustomFields, error) {
	path := "/customfields"
	response, _ := Get(config, path, nil)
	return resp.CustomFieldsFromJSON(response)
}

func QueryCustomFieldsByIds(config Config, pathId []string) (resp.CustomFields, error) {
	path := fmt.Sprintf("/customfields/%s", strings.Join(pathId, ","))
	response, _ := Get(config, path, nil)
	return resp.CustomFieldsFromJSON(response)
}

func CreateCustomFields(config Config, params params.CreateCustomFields) (resp.CustomFields, error) {
	path := "/customfields"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.CustomFieldsFromJSON(response)
}

func ModifyCustomFieldsById(config Config, params params.ModifyCustomFields, pathId string) (resp.CustomFields, error) {
	path := fmt.Sprintf("/customfields/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.CustomFieldsFromJSON(response)
}