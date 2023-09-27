package wrikego

import (
	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryColors(config Config) (resp.Colors, error) {
	path := "/colors"
	response, _ := Get(config, path, nil)
	return resp.ColorsFromJSON(response)
}