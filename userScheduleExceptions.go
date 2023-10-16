package wrikego

import (
	"fmt"
	

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryUserScheduleExceptions(config Config, params params.QueryUserScheduleExceptions) (resp.UserScheduleExceptions, error) {
	path := "/user_schedule_exclusions"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.UserScheduleExceptionsFromJSON(response)
}

func QueryUserScheduleExceptionsById(config Config, pathId string) (resp.UserScheduleExceptions, error) {
	path := fmt.Sprintf("/user_schedule_exclusions/%s", pathId)
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.UserScheduleExceptionsFromJSON(response)
}

func CreateUserScheduleExceptions(config Config, params params.CreateUserScheduleExceptions) (resp.UserScheduleExceptions, error) {
	path := "/user_schedule_exclusions"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.UserScheduleExceptionsFromJSON(response)
}

func ModifyUserScheduleExceptionsById(config Config, params params.ModifyUserScheduleExceptions) (resp.UserScheduleExceptions, error) {
	path := "/user_schedule_exclusions"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.UserScheduleExceptionsFromJSON(response)
}

func DeleteUserScheduleExceptionsById(config Config, pathId string) (resp.UserScheduleExceptions, error) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	response, err := Delete(config, path, nil)
	ErrorCheck(err)
	return resp.UserScheduleExceptionsFromJSON(response)
}