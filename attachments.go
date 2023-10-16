package wrikego

import (
	"fmt"
	
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryAttachments(config Config, params params.QueryAttachments) (resp.Attachments, error) {
	path := "/attachments"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.AttachmentsFromJSON(response)
}

func QueryAttachmentsByFolder(config Config, params params.QueryAttachments, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/folders/%s/attachments", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.AttachmentsFromJSON(response)
}

func QueryAttachmentsByTask(config Config, params params.QueryAttachments, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/tasks/%s/attachments", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.AttachmentsFromJSON(response)
}

func QueryAttachmentsByIds(config Config, params params.QueryAttachments, pathId []string) (resp.Attachments, error) {
	path := fmt.Sprintf("/attachments/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Get(config, path, body)
	ErrorCheck(err)
	return resp.AttachmentsFromJSON(response)
}

func DownloadAttachmentsByAttachment(config Config, pathId string) ([]byte, error) {
	path := fmt.Sprintf("/attachments/%s/download", pathId)
	response, err := Download(config, path, nil)
	ErrorCheck(err)
	return response, err
}

func DownloadAttachmentsPreviewByAttachment(config Config, params params.DownloadAttachmentPreview, pathId string) ([]byte, error) {
	path := fmt.Sprintf("/attachments/%s/preview", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, err := Download(config, path, body)
	ErrorCheck(err)
	return response, err
}

func QueryAttachmentsAccessUrlByAttachment(config Config, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/attachments/%s/url", pathId)
	response, err := Get(config, path, nil)
	ErrorCheck(err)
	return resp.AttachmentsFromJSON(response)
}

func UploadAttachmentsByFolder(config Config, params params.UploadAttachment, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/folders/%s/attachments", pathId)
	response, err := Upload(config, path, params)
	ErrorCheck(err)
	return resp.AttachmentsFromJSON(response)
}

func UploadAttachmentsByTask(config Config, params params.UploadAttachment, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/tasks/%s/attachments", pathId)
	response, err := Upload(config, path, params)
	ErrorCheck(err)
	return resp.AttachmentsFromJSON(response)
}

func UpdateAttachmentsById(config Config, params params.UploadAttachment, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/tasks/%s/attachments", pathId)
	response, err := Update(config, path, params)
	ErrorCheck(err)
	return resp.AttachmentsFromJSON(response)
}

func DeleteAttachmentsById(config Config, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/attachments/%s", pathId)
	response, err := Delete(config, path, nil)
	ErrorCheck(err)
	return resp.AttachmentsFromJSON(response)
}