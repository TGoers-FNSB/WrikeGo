package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryWebhooks(config Config) (resp.Webhooks, *http.Response) {
	path := "/webhooks"
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.WebhooksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryWebhooksByIds(config Config, pathId []string) (resp.Webhooks, *http.Response) {
	path := fmt.Sprintf("/webhooks/%s", strings.Join(pathId, ","))
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.WebhooksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateWebhooksByFolder(config Config, params params.CreateWebooks, pathId string) (resp.Webhooks, *http.Response) {
	path := fmt.Sprintf("/folders/%s/webhooks", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.WebhooksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateWebhooksBySpace(config Config, params params.CreateWebooks, pathId string) (resp.Webhooks, *http.Response) {
	path := fmt.Sprintf("/spaces/%s/webhooks", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.WebhooksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateWebhooks(config Config, params params.CreateWebooks) (resp.Webhooks, *http.Response) {
	path := "/webhooks"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.WebhooksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyWebhooksById(config Config, params params.ModifyWebhooks, pathId string) (resp.Webhooks, *http.Response) {
	path := fmt.Sprintf("/webhooks/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.WebhooksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteWebhooksById(config Config, pathId string) (resp.Webhooks, *http.Response) {
	path := fmt.Sprintf("/webhooks/%s", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.WebhooksFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}