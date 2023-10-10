package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func ModifyHourlyRatesProvisionsByIds(config Config, params params.ModifyHourlyRatesProvision, pathId []string) (resp.HourlyRatesProvision, error) {
	path := fmt.Sprintf("/contacts/%s/hourly_rates_provision", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.HourlyRatesProvisionFromJSON(response)
}