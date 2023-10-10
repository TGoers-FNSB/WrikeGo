package wrikego

import (
	"fmt"
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryUserScheduleExceptions(config Config, params params.QueryUserScheduleExceptions) (resp.UserScheduleExceptions, error) {
	path := "/user_schedule_exclusions"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.UserScheduleExceptionsFromJSON(response)
}

func QueryUserScheduleExceptionsById(config Config, pathId string) (resp.UserScheduleExceptions, error) {
	path := fmt.Sprintf("/user_schedule_exclusions/%s", pathId)
	response, _ := Get(config, path, nil)
	return resp.UserScheduleExceptionsFromJSON(response)
}

func CreateUserScheduleExceptions(config Config, params params.CreateUserScheduleExceptions) (resp.UserScheduleExceptions, error) {
	path := "/user_schedule_exclusions"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.UserScheduleExceptionsFromJSON(response)
}

func ModifyUserScheduleExceptionsById(config Config, params params.ModifyUserScheduleExceptions) (resp.UserScheduleExceptions, error) {
	path := "/user_schedule_exclusions"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.UserScheduleExceptionsFromJSON(response)
}

func DeleteUserScheduleExceptionsById(config Config, pathId string) (resp.UserScheduleExceptions, error) {
	path := fmt.Sprintf("/workschedules/%s/workschedule_exclusions", pathId)
	response, _ := Delete(config, path, nil)
	return resp.UserScheduleExceptionsFromJSON(response)
}