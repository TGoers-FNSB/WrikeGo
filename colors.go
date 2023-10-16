package wrikego

import (
	"net/http"

	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryColors(config Config) (resp.Colors, *http.Response) {
	path := "/colors"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.ColorsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}