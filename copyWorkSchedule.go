package wrikego

import (
	"fmt"
	

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func CreateWorkSchedulesByWorkSchedule(config Config, params params.CreateCopyWorkSchedules, pathId string) (resp.CopyWorkSchedules, error) {
	path := fmt.Sprintf("/workschedules/%s/duplicate", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.CopyWorkSchedulesFromJSON(response)
}