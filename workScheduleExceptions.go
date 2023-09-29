package wrikego

import (
	"fmt"
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryWorkScheduleExceptionsById(config Config, pathId string) (resp.WorkScheduleExceptions, error) {
	path := fmt.Sprintf("/workschedule_exclusions/%s", pathId)
	response, _ := Get(config, path, nil)
	return resp.WorkScheduleExceptionsFromJSON(response)
}

func QueryWorkScheduleExceptionsByWorkSchedule(config Config, params params.QueryWorkScheduleExceptions, pathId string) (resp.WorkScheduleExceptions, error) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.WorkScheduleExceptionsFromJSON(response)
}

func CreateWorkScheduleExceptionsByWorkSchedule(config Config, params params.CreateWorkScheduleExceptions, pathId string) (resp.WorkScheduleExceptions, error) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.WorkScheduleExceptionsFromJSON(response)
}

func ModifyWorkScheduleExceptionsByWorkSchedule(config Config, params params.ModifyWorkScheduleExceptions, pathId string) (resp.WorkScheduleExceptions, error) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.WorkScheduleExceptionsFromJSON(response)
}

func DeleteWorkScheduleExceptionsById(config Config, pathId string) (resp.WorkScheduleExceptions, error) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	response, _ := Delete(config, path, nil)
	return resp.WorkScheduleExceptionsFromJSON(response)
}