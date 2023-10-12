package wrikego

import (
	"log"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryAuditLog(config Config, params params.QueryAuditLog) (resp.AuditLog, error) {
	path := "/audit_log"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.AuditLogFromJSON(response)
}