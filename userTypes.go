package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryUserTypes(config Config) (resp.UserTypes, error) {
	path := "/user_types"
	response, _ := Get(config, path, nil)
	return resp.UserTypesFromJSON(response)
}