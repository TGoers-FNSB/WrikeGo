package wrikego

import (
	"fmt"
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryDataExport(config Config) (resp.DataExport, *http.Response) {
	path := "/data_export"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.DataExportFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryDataExportById(config Config, pathId string) (resp.DataExport, *http.Response) {
	path := fmt.Sprintf("/data_export/%s", pathId)
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.DataExportFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateDataExport(config Config) (resp.DataExport, *http.Response) {
	path := "/data_export"
	response, httpResponse, err := Post(config, path, nil)
	ErrorCheck(err)
	json, err := resp.DataExportFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryDataExportSchema(config Config, params params.QueryDataExportSchema) (resp.DataExport, *http.Response) {
	path := "/data_export_schema"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.DataExportFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}