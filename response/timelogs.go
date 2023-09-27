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
		Id          string `json:"id"`
		TaskId      string `json:"taskId"`
		UserId      string `json:"userId"`
		CategoryId  string `json:"categoryId,omitempty"`
		BillingType string `json:"billingType"`
		Hours       int    `json:"hours"`
		CreatedDate string `json:"createdDate"`
		UpdatedDate string `json:"updatedDate"`
		TrackedDate string `json:"trackedDate"`
		Comment     string `json:"comment"`
		Finance     struct {
			Currency   string  `json:"currency"`
			ActualFees float64 `json:"actualFees"`
			ActualCost float64 `json:"actualCost"`
		} `json:"finance"`
	} `json:"data"`
}
