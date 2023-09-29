package wrikeresponse

import (
	"encoding/json"
)

func WorkScheduleExceptionsFromJSON(data []byte) (WorkScheduleExceptions, error) {
	var item WorkScheduleExceptions
	err := json.Unmarshal(data, &item)
	return item, err
}

type WorkScheduleExceptions struct {
	Kind string `json:"kind"`
	Data []struct {
		Id            string `json:"id"`
		FromDate      string `json:"fromDate"`
		ToDate        string `json:"toDate"`
		IsWorkDays    bool   `json:"isWorkDays"`
		ExclusionType string `json:"exclusionType"`
	} `json:"data"`
}