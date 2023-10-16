package wrikego

import (
	"fmt"
	
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryHourlyRatesByContacts(config Config, pathId []string) (resp.HourlyRates, error) {
	path := fmt.Sprintf("/contacts/%s/hourly_rates", strings.Join(pathId, ","))
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.HourlyRatesFromJSON(response)
}

func QueryHourlyRatesByFolder(config Config, pathId string) (resp.HourlyRates, error) {
	path := fmt.Sprintf("/folders/%s/hourly_rates", pathId)
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.HourlyRatesFromJSON(response)
}

func QueryHourlyRatesByPlaceholders(config Config, pathId []string) (resp.HourlyRates, error) {
	path := fmt.Sprintf("/placeholders/%s/hourly_rates", strings.Join(pathId, ","))
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.HourlyRatesFromJSON(response)
}

func ModifyHourlyRates(config Config, params params.ModifyHourlyRates) (resp.HourlyRates, error) {
	path := "hourly_rates"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.HourlyRatesFromJSON(response)
}

func ModifyHourlyRatesByFolder(config Config, params params.ModifyHourlyRates, pathId string) (resp.HourlyRates, error) {
	path := fmt.Sprintf("/folders/%s/hourly_rates", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Put(config, path, body)
	ErrorCheck(err)
	return resp.HourlyRatesFromJSON(response)
}

func DeleteTeamMembersByFolder(config Config, params params.DeleteTeamMembers, pathId string) (resp.HourlyRates, error) {
	path := fmt.Sprintf("/folders/%s/project_team_members", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Delete(config, path, body)
	ErrorCheck(err)
	return resp.HourlyRatesFromJSON(response)
}