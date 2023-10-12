package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryWebhooks(config Config) (resp.Webhooks, error) {
	path := "/webhooks"
	response, _ := Get(config, path, nil)
	return resp.WebhooksFromJSON(response)
}

func QueryWebhooksByIds(config Config, pathId []string) (resp.Webhooks, error) {
	path := fmt.Sprintf("/webhooks/%s", strings.Join(pathId, ","))
	response, _ := Get(config, path, nil)
	return resp.WebhooksFromJSON(response)
}

func CreateWebhooksByFolder(config Config, params params.CreateWebooks, pathId string) (resp.Webhooks, error) {
	path := fmt.Sprintf("/folders/%s/webhooks", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.WebhooksFromJSON(response)
}

func CreateWebhooksBySpace(config Config, params params.CreateWebooks, pathId string) (resp.Webhooks, error) {
	path := fmt.Sprintf("/spaces/%s/webhooks", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.WebhooksFromJSON(response)
}

func CreateWebhooks(config Config, params params.CreateWebooks) (resp.Webhooks, error) {
	path := "/webhooks"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.WebhooksFromJSON(response)
}

func ModifyWebhooksById(config Config, params params.ModifyWebhooks, pathId string) (resp.Webhooks, error) {
	path := fmt.Sprintf("/webhooks/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.WebhooksFromJSON(response)
}

func DeleteWebhooksById(config Config, pathId string) (resp.Webhooks, error) {
	path := fmt.Sprintf("/webhooks/%s", pathId)
	response, _ := Delete(config, path, nil)
	return resp.WebhooksFromJSON(response)
}