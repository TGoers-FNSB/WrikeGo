package wrikego

import (
	"fmt"
	"strings"

	resp "github.com/TGoers-FNSB/WrikeGo/response"
)

func QueryPlaceholders(config Config) (resp.Placeholders, error) {
	path := "/placeholders"
	response, _ := Get(config, path, nil)
	return resp.PlaceholdersFromJSON(response)
}

func QueryPlaceholdersByIds(config Config, pathId []string) (resp.Placeholders, error) {
	path := fmt.Sprintf("/placeholders/%s", strings.Join(pathId, ","))
	response, _ := Get(config, path, nil)
	return resp.PlaceholdersFromJSON(response)
}
