package wrikego

import (
	"fmt"
	"net/url"

	types "github.com/TGoers-FNSB/WrikeGo/types"
)

func QueryTasks_Get_Tasks(config Config, params url.Values) (types.Tasks, error) {
	path := "/tasks"

	response, err := Get(config, path, params)
	fmt.Println(err)

	return types.TasksFromJSON(response)

}