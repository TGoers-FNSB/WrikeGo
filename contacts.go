package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryContacts(config Config, params params.QueryContacts) (resp.Contacts, error) {
	path := "/contacts"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.ContactsFromJSON(response)
}

func QueryContactsByIds(config Config, params params.QueryContacts, pathId []string) (resp.Contacts, error) {
	path := fmt.Sprintf("/contacts/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.ContactsFromJSON(response)
}

func QueryContactsFieldsHistoryByContacts(config Config, params params.QueryContacts, pathId []string) (resp.Contacts, error) {
	path := fmt.Sprintf("/contacts/%s/contacts_history", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.ContactsFromJSON(response)
}

func ModifyContactsById(config Config, params params.ModifyFolders, pathId string) (resp.Contacts, error) {
	path := fmt.Sprintf("/contacts/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.ContactsFromJSON(response)
}