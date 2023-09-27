package wrikego

import (
	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryTimelogCategories(config Config) (resp.Timelogs, error) {
	path := "/timelog_categories"
	response, _ := Get(config, path, nil)
	return resp.TimelogsFromJSON(response)
}