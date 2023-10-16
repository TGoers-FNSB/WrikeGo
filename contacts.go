package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryContacts(config Config, params params.QueryContacts) (resp.Contacts, *http.Response) {
	path := "/contacts"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.ContactsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryContactsByIds(config Config, params params.QueryContacts, pathId []string) (resp.Contacts, *http.Response) {
	path := fmt.Sprintf("/contacts/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.ContactsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryContactsFieldsHistoryByContacts(config Config, params params.QueryContacts, pathId []string) (resp.Contacts, *http.Response) {
	path := fmt.Sprintf("/contacts/%s/contacts_history", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.ContactsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyContactsById(config Config, params params.ModifyFolders, pathId string) (resp.Contacts, *http.Response) {
	path := fmt.Sprintf("/contacts/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.ContactsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}