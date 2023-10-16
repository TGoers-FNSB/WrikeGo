package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryTimelogs(config Config, params params.QueryTimelogs) (resp.Timelogs, *http.Response) {
	path := "/timelogs"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.TimelogsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryTimelogsByContact(config Config, params params.QueryTimelogs, pathId string) (resp.Timelogs, *http.Response) {
	path := fmt.Sprintf("/contacts/%s/timelogs", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.TimelogsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryTimelogsByFolders(config Config, params params.QueryTimelogs, pathId string) (resp.Timelogs, *http.Response) {
	path := fmt.Sprintf("/folders/%s/timelogs", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.TimelogsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryTimelogsByTask(config Config, params params.QueryTimelogs, pathId string) (resp.Timelogs, *http.Response) {
	path := fmt.Sprintf("/tasks/%s/timelogs", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.TimelogsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryTimelogsByTimelogCategory(config Config, params params.QueryTimelogs, pathId string) (resp.Timelogs, *http.Response) {
	path := fmt.Sprintf("/timelog_categories/%s/timelogs", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.TimelogsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryTimelogsByIds(config Config, params params.QueryTimelogs, pathId []string) (resp.Timelogs, *http.Response) {
	path := fmt.Sprintf("/timelogs/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.TimelogsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateTimelogsByTask(config Config, params params.CreateTimelogs, pathId string) (resp.Timelogs, *http.Response) {
	path := fmt.Sprintf("/tasks/%s/timelogs", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.TimelogsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyTimelogsById(config Config, params params.ModifyTimelogs, pathId string) (resp.Timelogs, *http.Response) {
	path := fmt.Sprintf("/timelogs/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.TimelogsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteTimelogsById(config Config, pathId string) (resp.Timelogs, *http.Response) {
	path := fmt.Sprintf("/timelogs/%s", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.TimelogsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}