package wrikego

import (
	"fmt"
	
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryComments(config Config, params params.QueryComments) (resp.Comments, error) {
	path := "/comments"
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.CommentsFromJSON(response)
}

func QueryCommentsByFolder(config Config, params params.QueryComments, pathId string) (resp.Comments, error) {
	path := fmt.Sprintf("/folders/%s/comments", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.CommentsFromJSON(response)
}

func QueryCommentsByTask(config Config, params params.QueryComments, pathId string) (resp.Comments, error) {
	path := fmt.Sprintf("/tasks/%s/comments", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.CommentsFromJSON(response)
}

func QueryCommentsByIds(config Config, params params.QueryComments, pathId []string) (resp.Comments, error) {
	path := fmt.Sprintf("/comments/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.CommentsFromJSON(response)
}

func CreateCommentsByFolder(config Config, params params.CreateComments, pathId string) (resp.Comments, error) {
	path := fmt.Sprintf("/folders/%s/comments", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.CommentsFromJSON(response)
}

func CreateCommentsByTask(config Config, params params.CreateComments, pathId string) (resp.Comments, error) {
	path := fmt.Sprintf("/folders/%s/comments", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Post(config, path, body)
	ErrorCheck(err)
	return resp.CommentsFromJSON(response)
}

func ModifyCommentsById(config Config, params params.ModifyComments, pathId string) (resp.Comments, error) {
	path := fmt.Sprintf("/comments/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.CommentsFromJSON(response)
}

func DeleteCommentById(config Config, pathId string) (resp.Comments, error) {
	path := fmt.Sprintf("/comments/%s", pathId)
	response, err := Delete(config, path, nil)
	ErrorCheck(err)
	return resp.CommentsFromJSON(response)
}