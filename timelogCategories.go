package wrikego

import (
	"net/http"

	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryTimelogCategories(config Config) (resp.TimelogCategories, *http.Response) {
	path := "/timelog_categories"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.TimelogCategoriesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}