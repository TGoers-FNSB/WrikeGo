package wrikego

import (
	"fmt"
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryFolderBlueprints(config Config) (resp.FolderBlueprints, *http.Response) {
	path := "/folder_blueprints"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.FolderBlueprintsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryFolderBlueprintsBySpace(config Config, pathId string) (resp.FolderBlueprints, *http.Response) {
	path := fmt.Sprintf("/spaces/%s/folder_blueprints", pathId)
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.FolderBlueprintsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateFolderBlueprintsAsync(config Config, params params.CreateFolderBlueprintsAsync, pathId string) (resp.FolderBlueprints, *http.Response) {
	path := fmt.Sprintf("/folder_blueprints/%s/launch_async", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.FolderBlueprintsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}