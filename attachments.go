package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/google/go-querystring/query"
)

func QueryAttachments(config Config, params params.QueryAttachments) (resp.Attachments, error) {
	path := "/attachments"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.AttachmentsFromJSON(response)
}

func QueryAttachmentsByFolder(config Config, params params.QueryAttachments, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/folders/%s/attachments", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.AttachmentsFromJSON(response)
}

func QueryAttachmentsByTask(config Config, params params.QueryAttachments, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/tasks/%s/attachments", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.AttachmentsFromJSON(response)
}

func QueryAttachmentsByIds(config Config, params params.QueryAttachments, pathId []string) (resp.Attachments, error) {
	path := fmt.Sprintf("/attachments/%s", strings.Join(pathId, ","))
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.AttachmentsFromJSON(response)
}

func DownloadAttachmentsByAttachment(config Config, pathId string) ([]byte, error) {
	path := fmt.Sprintf("/attachments/%s/download", pathId)
	response, err := Download(config, path, nil)
	return response, err
}

func DownloadAttachmentsPreviewByAttachment(config Config, params params.DownloadAttachmentPreview, pathId string) ([]byte, error) {
	path := fmt.Sprintf("/attachments/%s/preview", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, err := Download(config, path, body)
	return response, err
}

func QueryAttachmentsAccessUrlByAttachment(config Config, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/attachments/%s/url", pathId)
	response, _ := Get(config, path, nil)
	return resp.AttachmentsFromJSON(response)
}

func UploadAttachmentsByFolder(config Config, params params.UploadAttachment, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/folders/%s/attachments", pathId)
	response, _ := Upload(config, path, params)
	return resp.AttachmentsFromJSON(response)
}

func UploadAttachmentsByTask(config Config, params params.UploadAttachment, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/tasks/%s/attachments", pathId)
	response, _ := Upload(config, path, params)
	return resp.AttachmentsFromJSON(response)
}

func UpdateAttachmentsById(config Config, params params.UploadAttachment, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/tasks/%s/attachments", pathId)
	response, _ := Update(config, path, params)
	return resp.AttachmentsFromJSON(response)
}

func DeleteAttachmentsById(config Config, pathId string) (resp.Attachments, error) {
	path := fmt.Sprintf("/attachments/%s", pathId)
	response, _ := Delete(config, path, nil)
	return resp.AttachmentsFromJSON(response)
}