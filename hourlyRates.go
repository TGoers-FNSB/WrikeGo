package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryHourlyRatesByContacts(config Config, pathId []string) (resp.HourlyRates, *http.Response) {
	path := fmt.Sprintf("/contacts/%s/hourly_rates", strings.Join(pathId, ","))
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.HourlyRatesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryHourlyRatesByFolder(config Config, pathId string) (resp.HourlyRates, *http.Response) {
	path := fmt.Sprintf("/folders/%s/hourly_rates", pathId)
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.HourlyRatesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryHourlyRatesByPlaceholders(config Config, pathId []string) (resp.HourlyRates, *http.Response) {
	path := fmt.Sprintf("/placeholders/%s/hourly_rates", strings.Join(pathId, ","))
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.HourlyRatesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyHourlyRates(config Config, params params.ModifyHourlyRates) (resp.HourlyRates, *http.Response) {
	path := "hourly_rates"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.HourlyRatesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyHourlyRatesByFolder(config Config, params params.ModifyHourlyRates, pathId string) (resp.HourlyRates, *http.Response) {
	path := fmt.Sprintf("/folders/%s/hourly_rates", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.HourlyRatesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteTeamMembersByFolder(config Config, params params.DeleteTeamMembers, pathId string) (resp.HourlyRates, *http.Response) {
	path := fmt.Sprintf("/folders/%s/project_team_members", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Delete(config, path, body)
	ErrorCheck(err)
	json, err := resp.HourlyRatesFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}