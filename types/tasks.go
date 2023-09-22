package wriketypes

import (
	"encoding/json"
	"time"
)

func TasksFromJSON(data []byte) (Tasks, error) {
	var tasks Tasks
	err := json.Unmarshal(data, &tasks)
	return tasks, err
}

type Tasks struct {
	Kind string `json:"kind"`
	Data []struct {
		ID               string    `json:"id"`
		AccountID        string    `json:"accountId"`
		Title            string    `json:"title"`
		Description      string    `json:"description"`
		BriefDescription string    `json:"briefDescription"`
		ParentIds        []string  `json:"parentIds"`
		SuperParentIds   []string  `json:"superParentIds"`
		SharedIds        []string  `json:"sharedIds"`
		ResponsibleIds   []string  `json:"responsibleIds"`
		Status           string    `json:"status"`
		Importance       string    `json:"importance"`
		CreatedDate      time.Time `json:"createdDate"`
		UpdatedDate      time.Time `json:"updatedDate"`
		Dates            struct {
			Type     string `json:"type"`
			Duration int    `json:"duration"`
			Start    string `json:"start"`
			Due      string `json:"due"`
		} `json:"dates"`
		Scope           string   `json:"scope"`
		AuthorIds       []string `json:"authorIds"`
		CustomStatusID  string   `json:"customStatusId"`
		HasAttachments  bool     `json:"hasAttachments"`
		AttachmentCount int      `json:"attachmentCount"`
		Permalink       string   `json:"permalink"`
		Priority        string   `json:"priority"`
		SuperTaskIds    []string `json:"superTaskIds"`
		SubTaskIds      []string `json:"subTaskIds"`
		DependencyIds   []string `json:"dependencyIds"`
		Metadata        []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"metadata"`
		CustomFields []struct {
			ID    string `json:"id"`
			Value string `json:"value"`
		} `json:"customFields"`
	} `json:"data"`
}
