package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	types "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryFolders(config Config, params params.QueryFolders) (types.Folders, error) {
	path := "/folders"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return types.FoldersFromJSON(response)
}

func QueryFoldersInFolder(config Config, params params.QueryFolders, pathId string) (types.Folders, error) {
	path := fmt.Sprintf("/folders/%s/folders", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return types.FoldersFromJSON(response)
}

func QueryFoldersInSpace(config Config, params params.QueryFolders, pathId string) (types.Folders, error) {
	path := fmt.Sprintf("/spaces/%s/folders", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return types.FoldersFromJSON(response)
}

func QueryFoldersFieldsHistoryByIds(config Config, params params.QueryFolders, pathId []string) (types.Folders, error) {
	path := fmt.Sprintf("/folders/%s/folders_history", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return types.FoldersFromJSON(response)
}

func QueryFoldersByIds(config Config, params params.QueryFolders, pathId []string) (types.Folders, error) {
	path := fmt.Sprintf("/folders/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return types.FoldersFromJSON(response)
}

func CreateFolderInFolder(config Config, params params.CreateFolders, pathId string) (types.Folders, error) {
	path := fmt.Sprintf("/folders/%s/folders", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return types.FoldersFromJSON(response)
}

func CreateFolderCopyById(config Config, params params.CreateFoldersCopy, pathId string) (types.Folders, error) {
	path := fmt.Sprintf("/copy_folders/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return types.FoldersFromJSON(response)
}

func CreateFolderCopyAsyncById(config Config, params params.CreateFoldersCopy, pathId string) (types.Folders, error) {
	path := fmt.Sprintf("/copy_folders_async/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return types.FoldersFromJSON(response)
}

func ModifyFolderById(config Config, params params.ModifyFolders, pathId string) (types.Folders, error) {
	path := fmt.Sprintf("/folders/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return types.FoldersFromJSON(response)
}

func ModifyFoldersByIds(config Config, params params.ModifyFolders, pathId []string) (types.Folders, error) {
	path := fmt.Sprintf("/folders/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return types.FoldersFromJSON(response)
}

func DeleteFolderById(config Config, pathId string) (types.Tasks, error) {
	path := fmt.Sprintf("/folders/%s", pathId)
	response, _ := Delete(config, path)
	return types.TasksFromJSON(response)
}