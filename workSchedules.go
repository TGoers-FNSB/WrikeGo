package wrikego

import (
	"fmt"
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryWorkSchedules(config Config, params params.QueryWorkSchedules) (resp.WorkSchedules, *http.Response) {
	path := "/workschedules"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.WorkSchedulesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryWorkSchedulesById(config Config, params params.QueryWorkSchedules, pathId string) (resp.WorkSchedules, *http.Response) {
	path := fmt.Sprintf("/workschedules/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.WorkSchedulesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateWorkSchedules(config Config, params params.CreateWorkSchedules) (resp.WorkSchedules, *http.Response) {
	path := "/workschedules"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.WorkSchedulesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyWorkSchedulesById(config Config, params params.ModifyWorkSchedules, pathId string) (resp.WorkSchedules, *http.Response) {
	path := fmt.Sprintf("/workschedules/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.WorkSchedulesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteWorkSchedulesById(config Config, pathId string) (resp.WorkSchedules, *http.Response) {
	path := fmt.Sprintf("/workschedules/%s", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.WorkSchedulesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}
