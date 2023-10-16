package wrikego

import (
	"fmt"
	"net/http"
	"strings"

	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryPlaceholders(config Config) (resp.Placeholders, *http.Response) {
	path := "/placeholders"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.PlaceholdersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryPlaceholdersByIds(config Config, pathId []string) (resp.Placeholders, *http.Response) {
	path := fmt.Sprintf("/placeholders/%s", strings.Join(pathId, ","))
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.PlaceholdersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}
