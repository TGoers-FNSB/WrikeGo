package wrikego

import (
	"fmt"
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryWorkSchedules(config Config, params params.QueryWorkSchedules) (resp.WorkSchedules, error) {
	path := "/workschedules"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.WorkSchedulesFromJSON(response)
}

func QueryWorkSchedulesById(config Config, params params.QueryWorkSchedules, pathId string) (resp.WorkSchedules, error) {
	path := fmt.Sprintf("/workschedules/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.WorkSchedulesFromJSON(response)
}

func CreateWorkSchedules(config Config, params params.CreateWorkSchedules) (resp.WorkSchedules, error) {
	path := "/workschedules"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.WorkSchedulesFromJSON(response)
}

func ModifyWorkSchedulesById(config Config, params params.ModifyWorkSchedules, pathId string) (resp.WorkSchedules, error) {
	path := fmt.Sprintf("/workschedules/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.WorkSchedulesFromJSON(response)
}

func DeleteWorkSchedulesById(config Config, pathId string) (resp.WorkSchedules, error) {
	path := fmt.Sprintf("/workschedules/%s", pathId)
	response, _ := Delete(config, path, nil)
	return resp.WorkSchedulesFromJSON(response)
}
