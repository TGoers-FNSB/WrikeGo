package wrikego

import (
	"fmt"
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryFolderBlueprints(config Config) (resp.FolderBlueprints, error) {
	path := "/folder_blueprints"
	response, _ := Get(config, path, nil)
	return resp.FolderBlueprintsFromJSON(response)
}

func CreateFolderBlueprintsAsync(config Config, params params.CreateFolderBlueprintsAsync, pathId string) (resp.FolderBlueprints, error) {
	path := fmt.Sprintf("/folder_blueprints/%s/launch_async", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.FolderBlueprintsFromJSON(response)
}