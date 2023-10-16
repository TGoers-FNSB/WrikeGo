package wrikego

import (
	"fmt"
	

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryWorkSchedules(config Config, params params.QueryWorkSchedules) (resp.WorkSchedules, error) {
	path := "/workschedules"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.WorkSchedulesFromJSON(response)
}

func QueryWorkSchedulesById(config Config, params params.QueryWorkSchedules, pathId string) (resp.WorkSchedules, error) {
	path := fmt.Sprintf("/workschedules/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.WorkSchedulesFromJSON(response)
}

func CreateWorkSchedules(config Config, params params.CreateWorkSchedules) (resp.WorkSchedules, error) {
	path := "/workschedules"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.WorkSchedulesFromJSON(response)
}

func ModifyWorkSchedulesById(config Config, params params.ModifyWorkSchedules, pathId string) (resp.WorkSchedules, error) {
	path := fmt.Sprintf("/workschedules/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.WorkSchedulesFromJSON(response)
}

func DeleteWorkSchedulesById(config Config, pathId string) (resp.WorkSchedules, error) {
	path := fmt.Sprintf("/workschedules/%s", pathId)
	response, err := Delete(config, path, nil)
	ErrorCheck(err)
	return resp.WorkSchedulesFromJSON(response)
}
