package wrikego

import (
	"fmt"
	
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryContacts(config Config, params params.QueryContacts) (resp.Contacts, error) {
	path := "/contacts"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.ContactsFromJSON(response)
}

func QueryContactsByIds(config Config, params params.QueryContacts, pathId []string) (resp.Contacts, error) {
	path := fmt.Sprintf("/contacts/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.ContactsFromJSON(response)
}

func QueryContactsFieldsHistoryByContacts(config Config, params params.QueryContacts, pathId []string) (resp.Contacts, error) {
	path := fmt.Sprintf("/contacts/%s/contacts_history", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.ContactsFromJSON(response)
}

func ModifyContactsById(config Config, params params.ModifyFolders, pathId string) (resp.Contacts, error) {
	path := fmt.Sprintf("/contacts/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.ContactsFromJSON(response)
}