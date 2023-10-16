package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryComments(config Config, params params.QueryComments) (resp.Comments, *http.Response) {
	path := "/comments"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.CommentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryCommentsByFolder(config Config, params params.QueryComments, pathId string) (resp.Comments, *http.Response) {
	path := fmt.Sprintf("/folders/%s/comments", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.CommentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryCommentsByTask(config Config, params params.QueryComments, pathId string) (resp.Comments, *http.Response) {
	path := fmt.Sprintf("/tasks/%s/comments", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.CommentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryCommentsByIds(config Config, params params.QueryComments, pathId []string) (resp.Comments, *http.Response) {
	path := fmt.Sprintf("/comments/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.CommentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateCommentsByFolder(config Config, params params.CreateComments, pathId string) (resp.Comments, *http.Response) {
	path := fmt.Sprintf("/folders/%s/comments", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.CommentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateCommentsByTask(config Config, params params.CreateComments, pathId string) (resp.Comments, *http.Response) {
	path := fmt.Sprintf("/folders/%s/comments", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.CommentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyCommentsById(config Config, params params.ModifyComments, pathId string) (resp.Comments, *http.Response) {
	path := fmt.Sprintf("/comments/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.CommentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteCommentById(config Config, pathId string) (resp.Comments, *http.Response) {
	path := fmt.Sprintf("/comments/%s", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.CommentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}