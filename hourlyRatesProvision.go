package wrikego

import (
	"fmt"
	
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func ModifyHourlyRatesProvisionsByIds(config Config, params params.ModifyHourlyRatesProvision, pathId []string) (resp.HourlyRatesProvision, error) {
	path := fmt.Sprintf("/contacts/%s/hourly_rates_provision", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.HourlyRatesProvisionFromJSON(response)
}