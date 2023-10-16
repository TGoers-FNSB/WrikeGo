package wrikego

import (
	"fmt"
	

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryDataExport(config Config) (resp.DataExport, error) {
	path := "/data_export"
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.DataExportFromJSON(response)
}

func QueryDataExportById(config Config, pathId string) (resp.DataExport, error) {
	path := fmt.Sprintf("/data_export/%s", pathId)
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.DataExportFromJSON(response)
}

func CreateDataExport(config Config) (resp.DataExport, error) {
	path := "/data_export"
	response, err := Post(config, path, nil)
	ErrorCheck(err)
	return resp.DataExportFromJSON(response)
}

func QueryDataExportSchema(config Config, params params.QueryDataExportSchema) (resp.DataExport, error) {
	path := "/data_export_schema"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.DataExportFromJSON(response)
}