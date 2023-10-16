package wrikego

import (
	"fmt"
	"net/http"

	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryAsyncJobById(config Config, pathId string) (resp.AsyncJob, *http.Response) {
	path := fmt.Sprintf("/async_job/%s", pathId)
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.AsyncJobFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}