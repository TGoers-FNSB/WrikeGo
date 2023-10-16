package wrikego

import (
	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryColors(config Config) (resp.Colors, error) {
	path := "/colors"
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.ColorsFromJSON(response)
}