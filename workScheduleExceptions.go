package wrikego

import (
	"fmt"
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryWorkScheduleExceptionsById(config Config, pathId string) (resp.WorkScheduleExceptions, *http.Response) {
	path := fmt.Sprintf("/workschedule_exclusions/%s", pathId)
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.WorkScheduleExceptionsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryWorkScheduleExceptionsByWorkSchedule(config Config, params params.QueryWorkScheduleExceptions, pathId string) (resp.WorkScheduleExceptions, *http.Response) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.WorkScheduleExceptionsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateWorkScheduleExceptionsByWorkSchedule(config Config, params params.CreateWorkScheduleExceptions, pathId string) (resp.WorkScheduleExceptions, *http.Response) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.WorkScheduleExceptionsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyWorkScheduleExceptionsById(config Config, params params.ModifyWorkScheduleExceptions, pathId string) (resp.WorkScheduleExceptions, *http.Response) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.WorkScheduleExceptionsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteWorkScheduleExceptionsById(config Config, pathId string) (resp.WorkScheduleExceptions, *http.Response) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.WorkScheduleExceptionsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}