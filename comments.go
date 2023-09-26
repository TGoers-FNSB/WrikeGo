package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryComments(config Config, params params.QueryComments) (resp.Comments, error) {
	path := "/comments"
	response, _ := Get(config, path, nil)
	return resp.CommentsFromJSON(response)
}

func QueryCommentsByFolder(config Config, params params.QueryComments, pathId string) (resp.Comments, error) {
	path := fmt.Sprintf("/folders/%s/comments", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.CommentsFromJSON(response)
}

func QueryCommentsByTask(config Config, params params.QueryComments, pathId string) (resp.Comments, error) {
	path := fmt.Sprintf("/tasks/%s/comments", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.CommentsFromJSON(response)
}

func QueryCommentsByIds(config Config, params params.QueryComments, pathId []string) (resp.Comments, error) {
	path := fmt.Sprintf("/comments/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.CommentsFromJSON(response)
}

func CreateCommentByFolder(config Config, params params.CreateComments, pathId string) (resp.Comments, error) {
	path := fmt.Sprintf("/folders/%s/comments", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.CommentsFromJSON(response)
}

func CreateCommentByTask(config Config, params params.CreateComments, pathId string) (resp.Comments, error) {
	path := fmt.Sprintf("/folders/%s/comments", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.CommentsFromJSON(response)
}

func ModifyCommentById(config Config, params params.ModifyComments, pathId string) (resp.Comments, error) {
	path := fmt.Sprintf("/comments/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.CommentsFromJSON(response)
}

func DeleteCommentById(config Config, pathId string) (resp.Comments, error) {
	path := fmt.Sprintf("/comments/%s", pathId)
	response, _ := Delete(config, path, nil)
	return resp.CommentsFromJSON(response)
}