package wrikeresponse

import (
	"encoding/json"
)

func TasksFromJSON(data []byte) (Tasks, error) {
	var item Tasks
	err := json.Unmarshal(data, &item)
	return item, err
}

type Tasks struct {
	Kind string `json:"kind,omitempty"`
	Data []struct {
		Id                      string   `json:"id,omitempty"`
		AccountId               string   `json:"accountId,omitempty"`
		Title                   string   `json:"title,omitempty"`
		Description             string   `json:"description,omitempty"`
		BriefDescription        string   `json:"briefDescription,omitempty"`
		ParentIds               []string `json:"parentIds,omitempty"`
		SuperParentIds          []string `json:"superParentIds,omitempty"`
		SharedIds               []string `json:"sharedIds,omitempty"`
		ResponsibleIds          []string `json:"responsibleIds,omitempty"`
		ResponsiblePlaceholders []string `json:"responsiblePlaceholderIds,omitempty"`
		Status                  string   `json:"status,omitempty"`
		Importance              string   `json:"importance,omitempty"`
		CreatedDate             string   `json:"createdDate,omitempty"`
		UpdatedDate             string   `json:"updatedDate,omitempty"`
		CompletedDate           string   `json:"completedDate,omitempty"`
		Dates                   struct {
			Type           string `json:"type,omitempty"`
			Duration       int    `json:"duration,omitempty"`
			Start          string `json:"start,omitempty"`
			Due            string `json:"due,omitempty"`
			WorkOnWeekends bool   `json:"workOnWeekends,omitempty"`
		} `json:"dates,omitempty"`
		Scope           string   `json:"scope,omitempty"`
		AuthorIds       []string `json:"authorIds,omitempty"`
		CustomStatusId  string   `json:"customStatusId,omitempty"`
		HasAttachments  bool     `json:"hasAttachments,omitempty"`
		AttachmentCount int      `json:"attachmentCount,omitempty"`
		Permalink       string   `json:"permalink,omitempty"`
		Priority        string   `json:"priority,omitempty"`
		FollowedByMe    bool     `json:"followedByMe,omitempty"`
		FollowerIds     []string `json:"followerIds,omitempty"`
		Recurrent       bool     `json:"recurrent,omitempty"`
		SuperTaskIds    []string `json:"superTaskIds,omitempty"`
		SubTaskIds      []string `json:"subTaskIds,omitempty"`
		DependencyIds   []string `json:"dependencyIds,omitempty"`
		Metadata        []struct {
			Key   string `json:"key,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"metadata,omitempty"`
		CustomFields []struct {
			ID    string `json:"id,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"customFields,omitempty"`
		BillingType      string `json:"billingType,omitempty"`
		EffortAllocation struct {
			AllocatedEffort           int     `json:"allocatedEffort,omitempty"`
			DailyAllocationPercentage float64 `json:"dailyAllocationPercentage,omitempty"`
			Mode                      string  `json:"mode,omitempty"`
			TotalEffort               int     `json:"totalEffort,omitempty"`
		}
		Finance struct {
			Currency    string  `json:"currency,omitempty"`
			ActualFees  float64 `json:"actualFees,omitempty"`
			ActualCost  float64 `json:"actualCost,omitempty"`
			PlannedFees float64 `json:"plannedFees,omitempty"`
			PlannedCost float64 `json:"plannedCost,omitempty"`
		} `json:"finance,omitempty"`
		CustomItemTypeID string  `json:"customItemTypeId,omitempty"`
		ActualFees       float64 `json:"actualFees,omitempty"`
		ActualCost       float64 `json:"actualCost,omitempty"`
		PlannedFees      float64 `json:"plannedFees,omitempty"`
		PlannedCost      float64 `json:"plannedCost,omitempty"`
	} `json:"data,omitempty"`
}
