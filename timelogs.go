package wrikego

import (
	"fmt"
	
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryTimelogs(config Config, params params.QueryTimelogs) (resp.Timelogs, error) {
	path := "/timelogs"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.TimelogsFromJSON(response)
}

func QueryTimelogsByContact(config Config, params params.QueryTimelogs, pathId string) (resp.Timelogs, error) {
	path := fmt.Sprintf("/contacts/%s/timelogs", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.TimelogsFromJSON(response)
}

func QueryTimelogsByFolders(config Config, params params.QueryTimelogs, pathId string) (resp.Timelogs, error) {
	path := fmt.Sprintf("/folders/%s/timelogs", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.TimelogsFromJSON(response)
}

func QueryTimelogsByTask(config Config, params params.QueryTimelogs, pathId string) (resp.Timelogs, error) {
	path := fmt.Sprintf("/tasks/%s/timelogs", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.TimelogsFromJSON(response)
}

func QueryTimelogsByTimelogCategory(config Config, params params.QueryTimelogs, pathId string) (resp.Timelogs, error) {
	path := fmt.Sprintf("/timelog_categories/%s/timelogs", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.TimelogsFromJSON(response)
}

func QueryTimelogsByIds(config Config, params params.QueryTimelogs, pathId []string) (resp.Timelogs, error) {
	path := fmt.Sprintf("/timelogs/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.TimelogsFromJSON(response)
}

func CreateTimelogsByTask(config Config, params params.CreateTimelogs, pathId string) (resp.Timelogs, error) {
	path := fmt.Sprintf("/tasks/%s/timelogs", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.TimelogsFromJSON(response)
}

func ModifyTimelogsById(config Config, params params.ModifyTimelogs, pathId string) (resp.Timelogs, error) {
	path := fmt.Sprintf("/timelogs/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.TimelogsFromJSON(response)
}

func DeleteTimelogsById(config Config, pathId string) (resp.Timelogs, error) {
	path := fmt.Sprintf("/timelogs/%s", pathId)
	response, err := Delete(config, path, nil)
	ErrorCheck(err)
	return resp.TimelogsFromJSON(response)
}