package wrikego

import (
	"fmt"
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func CreateWorkSchedulesByWorkSchedule(config Config, params params.CreateCopyWorkSchedules, pathId string) (resp.CopyWorkSchedules, *http.Response) {
	path := fmt.Sprintf("/workschedules/%s/duplicate", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.CopyWorkSchedulesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}