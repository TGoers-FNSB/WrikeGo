package wrikeresponse

import "encoding/json"

func WebhooksFromJSON(data []byte) (Webhooks, error) {
	var item Webhooks
	err := json.Unmarshal(data, &item)
	return item, err
}

type Webhooks []struct {
	TaskID          string    `json:"taskId"`
	WebhookID       string    `json:"webhookId"`
	EventAuthorID   string    `json:"eventAuthorId"`
	EventType       string    `json:"eventType"`
	LastUpdatedDate string `json:"lastUpdatedDate"`

	// TaskStatusChanged
	OldStatus         string `json:"oldStatus"`
	Status            string `json:"status"`
	OldCustomStatusID string `json:"oldCustomStatusId"`
	CustomStatusID    string `json:"customStatusId"`

	// CommentAdded
	CommentID string `json:"commentId"`
	Comment   struct {
		Text string `json:"text"`
		HTML string `json:"html"`
	} `json:"comment"`
}