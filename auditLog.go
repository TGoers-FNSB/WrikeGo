package wrikego

import (
	

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryAuditLog(config Config, params params.QueryAuditLog) (resp.AuditLog, error) {
	path := "/audit_log"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.AuditLogFromJSON(response)
}