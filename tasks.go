package wrikego

import (
	"net/url"

	types "github.com/TGoers-FNSB/WrikeGo/types"
)

func QueryTasks_Get_Tasks(config Config, params url.Values) (types.Tasks, error) {
	path := "/tasks"

	response, _ := Get(config, path, params)

	return types.TasksFromJSON(response)
	// Comments5
}
