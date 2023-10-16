package wrikego

import (
	"net/http"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryAuditLog(config Config, params params.QueryAuditLog) (resp.AuditLog, *http.Response) {
	path := "/audit_log"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.AuditLogFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}