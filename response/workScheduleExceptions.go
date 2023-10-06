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
		Id            *string `json:"id,omitempty"`
		FromDate      *string `json:"fromDate,omitempty"`
		ToDate        *string `json:"toDate,omitempty"`
		IsWorkDays    *bool   `json:"isWorkDays,omitempty"`
		ExclusionType *string `json:"exclusionType,omitempty"`
	} `json:"data"`
}