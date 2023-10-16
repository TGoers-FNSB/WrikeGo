package wrikego

import (
	"fmt"

	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryAsyncJobById(config Config, pathId string) (resp.AsyncJob, error) {
	path := fmt.Sprintf("/async_job/%s", pathId)
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.AsyncJobFromJSON(response)
}