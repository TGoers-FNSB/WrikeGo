package wrikego

import (
	"fmt"
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryDataExport(config Config) (resp.DataExport, error) {
	path := "/data_export"
	response, _ := Get(config, path, nil)
	return resp.DataExportFromJSON(response)
}

func QueryDataExportById(config Config, pathId string) (resp.DataExport, error) {
	path := fmt.Sprintf("/data_export/%s", pathId)
	response, _ := Get(config, path, nil)
	return resp.DataExportFromJSON(response)
}

func CreateDataExport(config Config) (resp.DataExport, error) {
	path := "/data_export"
	response, _ := Post(config, path, nil)
	return resp.DataExportFromJSON(response)
}

func QueryDataExportSchema(config Config, params params.QueryDataExportSchema) (resp.DataExport, error) {
	path := "/data_export_schema"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.DataExportFromJSON(response)
}