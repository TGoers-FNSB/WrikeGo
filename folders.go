package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryFolders(config Config, params params.QueryFolders) (resp.Folders, error) {
	path := "/folders"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.FoldersFromJSON(response)
}

func QueryFoldersByFolder(config Config, params params.QueryFolders, pathId string) (resp.Folders, error) {
	path := fmt.Sprintf("/folders/%s/folders", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.FoldersFromJSON(response)
}

func QueryFoldersBySpace(config Config, params params.QueryFolders, pathId string) (resp.Folders, error) {
	path := fmt.Sprintf("/spaces/%s/folders", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.FoldersFromJSON(response)
}

func QueryFoldersFieldsHistoryByFolders(config Config, params params.QueryFolders, pathId []string) (resp.Folders, error) {
	path := fmt.Sprintf("/folders/%s/folders_history", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.FoldersFromJSON(response)
}

func QueryFoldersByIds(config Config, params params.QueryFolders, pathId []string) (resp.Folders, error) {
	path := fmt.Sprintf("/folders/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.FoldersFromJSON(response)
}

func CreateFoldersByFolder(config Config, params params.CreateFolders, pathId string) (resp.Folders, error) {
	path := fmt.Sprintf("/folders/%s/folders", pathId)
	body, err := Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.FoldersFromJSON(response)
}

func CreateFolderCopyById(config Config, params params.CreateFoldersCopy, pathId string) (resp.Folders, error) {
	path := fmt.Sprintf("/copy_folders/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.FoldersFromJSON(response)
}

func CreateFolderCopyAsyncById(config Config, params params.CreateFoldersCopy, pathId string) (resp.Folders, error) {
	path := fmt.Sprintf("/copy_folders_async/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.FoldersFromJSON(response)
}

func ModifyFoldersById(config Config, params params.ModifyFolders, pathId string) (resp.Folders, error) {
	path := fmt.Sprintf("/folders/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.FoldersFromJSON(response)
}

func ModifyFoldersByIds(config Config, params params.ModifyFolders, pathId []string) (resp.Folders, error) {
	path := fmt.Sprintf("/folders/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.FoldersFromJSON(response)
}

func DeleteFoldersById(config Config, pathId string) (resp.Tasks, error) {
	path := fmt.Sprintf("/folders/%s", pathId)
	response, _ := Delete(config, path, nil)
	return resp.TasksFromJSON(response)
}