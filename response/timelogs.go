package wrikeresponse

import "encoding/json"

func TimelogsFromJSON(data []byte) (Timelogs, error) {
	var item Timelogs
	err := json.Unmarshal(data, &item)
	return item, err
}

type Timelogs struct {
	Kind string `json:"kind"`
	Data []struct {
		Id          *string `json:"id,omitempty"`
		TaskId      *string `json:"taskId,omitempty"`
		UserId      *string `json:"userId,omitempty"`
		CategoryId  *string `json:"categoryId,omitempty"`
		BillingType *string `json:"billingType,omitempty"`
		Hours       *int    `json:"hours,omitempty"`
		CreatedDate *string `json:"createdDate,omitempty"`
		UpdatedDate *string `json:"updatedDate,omitempty"`
		TrackedDate *string `json:"trackedDate,omitempty"`
		Comment     *string `json:"comment,omitempty"`
		Finance     *struct {
			Currency   *string  `json:"currency,omitempty"`
			ActualFees *float64 `json:"actualFees,omitempty"`
			ActualCost *float64 `json:"actualCost,omitempty"`
		} `json:"finance,omitempty"`
	} `json:"data"`
}
