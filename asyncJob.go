package wrikego

import (
	"fmt"

	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryAsyncJobById(config Config, pathId string) (resp.AsyncJob, error) {
	path := fmt.Sprintf("/async_job/%s", pathId)
	response, _ := Get(config, path, nil)
	return resp.AsyncJobFromJSON(response)
}