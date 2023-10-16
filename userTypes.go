package wrikego

import (
	"net/http"

	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryUserTypes(config Config) (resp.UserTypes, *http.Response) {
	path := "/user_types"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.UserTypesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}
