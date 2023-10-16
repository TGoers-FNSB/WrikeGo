package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryFolders(config Config, params params.QueryFolders) (resp.Folders, *http.Response) {
	path := "/folders"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.FoldersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryFoldersByFolder(config Config, params params.QueryFolders, pathId string) (resp.Folders, *http.Response) {
	path := fmt.Sprintf("/folders/%s/folders", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.FoldersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryFoldersBySpace(config Config, params params.QueryFolders, pathId string) (resp.Folders, *http.Response) {
	path := fmt.Sprintf("/spaces/%s/folders", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.FoldersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryFoldersFieldsHistoryByFolders(config Config, params params.QueryFolders, pathId []string) (resp.Folders, *http.Response) {
	path := fmt.Sprintf("/folders/%s/folders_history", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.FoldersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryFoldersByIds(config Config, params params.QueryFolders, pathId []string) (resp.Folders, *http.Response) {
	path := fmt.Sprintf("/folders/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.FoldersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateFoldersByFolder(config Config, params params.CreateFolders, pathId string) (resp.Folders, *http.Response) {
	path := fmt.Sprintf("/folders/%s/folders", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.FoldersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateFolderCopyById(config Config, params params.CreateFoldersCopy, pathId string) (resp.Folders, *http.Response) {
	path := fmt.Sprintf("/copy_folders/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.FoldersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateFolderCopyAsyncById(config Config, params params.CreateFoldersCopy, pathId string) (resp.Folders, *http.Response) {
	path := fmt.Sprintf("/copy_folders_async/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.FoldersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyFoldersById(config Config, params params.ModifyFolders, pathId string) (resp.Folders, *http.Response) {
	path := fmt.Sprintf("/folders/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.FoldersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyFoldersByIds(config Config, params params.ModifyFolders, pathId []string) (resp.Folders, *http.Response) {
	path := fmt.Sprintf("/folders/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	json, err := resp.FoldersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteFoldersById(config Config, pathId string) (resp.Folders, *http.Response) {
	path := fmt.Sprintf("/folders/%s", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.FoldersFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}