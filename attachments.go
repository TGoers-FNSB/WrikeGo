package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryAttachments(config Config, params params.QueryAttachments) (resp.Attachments, *http.Response) {
	path := "/attachments"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.AttachmentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryAttachmentsByFolder(config Config, params params.QueryAttachments, pathId string) (resp.Attachments, *http.Response) {
	path := fmt.Sprintf("/folders/%s/attachments", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.AttachmentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryAttachmentsByTask(config Config, params params.QueryAttachments, pathId string) (resp.Attachments, *http.Response) {
	path := fmt.Sprintf("/tasks/%s/attachments", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.AttachmentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryAttachmentsByIds(config Config, params params.QueryAttachments, pathId []string) (resp.Attachments, *http.Response) {
	path := fmt.Sprintf("/attachments/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.AttachmentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DownloadAttachmentsByAttachment(config Config, pathId string) ([]byte, *http.Response) {
	path := fmt.Sprintf("/attachments/%s/download", pathId)
	response, httpResponse, err := Download(config, path, nil)
	ErrorCheck(err)
	return response, httpResponse
}

func DownloadAttachmentsPreviewByAttachment(config Config, params params.DownloadAttachmentPreview, pathId string) ([]byte, *http.Response) {
	path := fmt.Sprintf("/attachments/%s/preview", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Download(config, path, body)
	ErrorCheck(err)
	return response, httpResponse
}

func QueryAttachmentsAccessUrlByAttachment(config Config, pathId string) (resp.Attachments, *http.Response) {
	path := fmt.Sprintf("/attachments/%s/url", pathId)
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.AttachmentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func UploadAttachmentsByFolder(config Config, params params.UploadAttachment, pathId string) (resp.Attachments, *http.Response) {
	path := fmt.Sprintf("/folders/%s/attachments", pathId)
	response, httpResponse, err := Upload(config, path, params)
	ErrorCheck(err)
	json, err := resp.AttachmentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func UploadAttachmentsByTask(config Config, params params.UploadAttachment, pathId string) (resp.Attachments, *http.Response) {
	path := fmt.Sprintf("/tasks/%s/attachments", pathId)
	response, httpResponse, err := Upload(config, path, params)
	ErrorCheck(err)
	json, err := resp.AttachmentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func UpdateAttachmentsById(config Config, params params.UploadAttachment, pathId string) (resp.Attachments, *http.Response) {
	path := fmt.Sprintf("/tasks/%s/attachments", pathId)
	response, httpResponse, err := Update(config, path, params)
	ErrorCheck(err)
	json, err := resp.AttachmentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteAttachmentsById(config Config, pathId string) (resp.Attachments, *http.Response) {
	path := fmt.Sprintf("/attachments/%s", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.AttachmentsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}