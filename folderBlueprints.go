package wrikego

import (
	"fmt"
	

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryFolderBlueprints(config Config) (resp.FolderBlueprints, error) {
	path := "/folder_blueprints"
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.FolderBlueprintsFromJSON(response)
}

func QueryFolderBlueprintsBySpace(config Config, pathId string) (resp.FolderBlueprints, error) {
	path := fmt.Sprintf("/spaces/%s/folder_blueprints", pathId)
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.FolderBlueprintsFromJSON(response)
}

func CreateFolderBlueprintsAsync(config Config, params params.CreateFolderBlueprintsAsync, pathId string) (resp.FolderBlueprints, error) {
	path := fmt.Sprintf("/folder_blueprints/%s/launch_async", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.FolderBlueprintsFromJSON(response)
}