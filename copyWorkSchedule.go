package wrikego

import (
	"fmt"
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func CreateWorkSchedulesByWorkScheduleId(config Config, params params.CreateCopyWorkSchedules, pathId string) (resp.CopyWorkSchedules, error) {
	path := fmt.Sprintf("/workschedules/%s/duplicate", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.CopyWorkSchedulesFromJSON(response)
}