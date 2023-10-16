package wrikego

import (
	"fmt"
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryUserScheduleExceptions(config Config, params params.QueryUserScheduleExceptions) (resp.UserScheduleExceptions, *http.Response) {
	path := "/user_schedule_exclusions"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.UserScheduleExceptionsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryUserScheduleExceptionsById(config Config, pathId string) (resp.UserScheduleExceptions, *http.Response) {
	path := fmt.Sprintf("/user_schedule_exclusions/%s", pathId)
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.UserScheduleExceptionsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateUserScheduleExceptions(config Config, params params.CreateUserScheduleExceptions) (resp.UserScheduleExceptions, *http.Response) {
	path := "/user_schedule_exclusions"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.UserScheduleExceptionsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyUserScheduleExceptionsById(config Config, params params.ModifyUserScheduleExceptions) (resp.UserScheduleExceptions, *http.Response) {
	path := "/user_schedule_exclusions"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.UserScheduleExceptionsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteUserScheduleExceptionsById(config Config, pathId string) (resp.UserScheduleExceptions, *http.Response) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.UserScheduleExceptionsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}