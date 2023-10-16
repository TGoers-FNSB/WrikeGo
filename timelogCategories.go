package wrikego

import (
	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryTimelogCategories(config Config) (resp.Timelogs, error) {
	path := "/timelog_categories"
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.TimelogsFromJSON(response)
}