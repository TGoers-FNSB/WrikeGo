package wrikego

import (
	"fmt"
	

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryWorkScheduleExceptionsById(config Config, pathId string) (resp.WorkScheduleExceptions, error) {
	path := fmt.Sprintf("/workschedule_exclusions/%s", pathId)
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.WorkScheduleExceptionsFromJSON(response)
}

func QueryWorkScheduleExceptionsByWorkSchedule(config Config, params params.QueryWorkScheduleExceptions, pathId string) (resp.WorkScheduleExceptions, error) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.WorkScheduleExceptionsFromJSON(response)
}

func CreateWorkScheduleExceptionsByWorkSchedule(config Config, params params.CreateWorkScheduleExceptions, pathId string) (resp.WorkScheduleExceptions, error) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.WorkScheduleExceptionsFromJSON(response)
}

func ModifyWorkScheduleExceptionsById(config Config, params params.ModifyWorkScheduleExceptions, pathId string) (resp.WorkScheduleExceptions, error) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.WorkScheduleExceptionsFromJSON(response)
}

func DeleteWorkScheduleExceptionsById(config Config, pathId string) (resp.WorkScheduleExceptions, error) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	response, err := Delete(config, path, nil)
	ErrorCheck(err)
	return resp.WorkScheduleExceptionsFromJSON(response)
}