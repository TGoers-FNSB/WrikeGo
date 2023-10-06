package wrikeresponse

import (
	"encoding/json"
)

func FoldersFromJSON(data []byte) (Folders, error) {
	var item Folders
	err := json.Unmarshal(data, &item)
	return item, err
}

type Folders struct {
	Kind string `json:"kind"`
	Data []struct {
		Id       *string   `json:"id,omitempty"`
		Title    *string   `json:"title,omitempty"`
		Color    *string   `json:"color,omitempty"`
		ChildIds *[]string `json:"childIds,omitempty"`
		Scope    *string   `json:"scope,omitempty"`
		Project  *struct {
			AuthorId       *string   `json:"authorId,omitempty"`
			OwnerIds       *[]string `json:"ownerIds,omitempty"`
			Status         *string   `json:"status,omitempty"`
			CustomStatusId *string   `json:"customStatusId,omitempty"`
			StartDate      *string   `json:"starDate,omitempty"`
			EndDate        *string   `json:"endDate,omitempty"`
			CreatedDate    *string   `json:"createdDate,omitempty"`
			CompletedDate  *string   `json:"completedDate,omitempty"`
			ContractType   *string   `json:"contractType,omitempty"`
			Finance        *struct {
				Currency    *string  `json:"currency,omitempty"`
				ActualFees  *float64 `json:"actualFees,omitempty"`
				ActualCost  *float64 `json:"actualCost,omitempty"`
				PlannedFees *float64 `json:"plannedFees,omitempty"`
				Budget      *float64 `json:"budget,omitempty"`
			} `json:"finance,omitempty"`
			Currency    *string  `json:"currency,omitempty"`
			ActualFees  *float64 `json:"actualFees,omitempty"`
			ActualCost  *float64 `json:"actualCost,omitempty"`
			PlannedFees *float64 `json:"plannedFees,omitempty"`
			Budget      *float64 `json:"budget,omitempty"`
		} `json:"project,omitempty"`
		Space            *bool     `json:"space,omitempty"`
		CustomItemTypeId *string   `json:"customItemTypeId,omitempty"`
		AccountId        *string   `json:"accountId,omitempty"`
		CreatedDate      *string    `json:"createdDate,omitempty"`
		UpdatedDate      *string     `json:"updatedDate,omitempty"`
		BriefDescription *string   `json:"briefDescription,omitempty"`
		Description      *string   `json:"description,omitempty"`
		SharedIds        *[]string `json:"sharedIds,omitempty"`
		ParentIds        *[]string `json:"parentIds,omitempty"`
		SuperParentIds   *[]string `json:"superParentIds,omitempty"`
		HasAttachments   *bool     `json:"hasAttachments,omitempty"`
		AttachmentCount  *int      `json:"attachmentCount,omitempty"`
		Permalink        *string   `json:"permalink,omitempty"`
		WorkflowId       *string   `json:"workflowId,omitempty"`
		Metadata         *[]struct {
			Key   *string `json:"key,omitempty"`
			Value *string `json:"value,omitempty"`
		} `json:"metadata,omitempty"`
		CustomFields *[]struct {
			Id    *string `json:"id,omitempty"`
			Value *string `json:"value,omitempty"`
		} `json:"customFields,omitempty"`
		CustomColumnIds *[]string `json:"customColumnIds,omitempty"`
		UserAccessRoles *struct {
			Key   *string `json:"key,omitempty"`
			Value *string `json:"value,omitempty"`
		} `json:"userAccessRoles,omitempty"`
		Status	*string	`josn:"status,omitempty"`
		ProgressPercent	*float64	`json:"progressPercent,omitempty"`
		ProcessedCount	*int	`json:"processedCount,omitempty"`
		Type	*string	`json:"type,omitempty"`
		Result	*interface{}	`json:"result,omitempty"` //? Unknown data type
		ErrorMessage	*string	`json:"errorMessage,omitempty"`
	} `json:"data"`
}
