package wrikeresponse

import (
	"encoding/json"
)

func UserScheduleExceptionsFromJSON(data []byte) (UserScheduleExceptions, error) {
	var item UserScheduleExceptions
	err := json.Unmarshal(data, &item)
	return item, err
}

type UserScheduleExceptions struct {
	Kind string `json:"kind"`
	Data []struct {
		Id            *string `json:"id,omitempty"`
		UserId        *string `json:"userId,omitempty"`
		FromDate      *string `json:"fromDate,omitempty"`
		ToDate        *string `json:"toDate,omitempty"`
		IsWorkDays    *bool   `json:"isWorkDays,omitempty"`
		ExclusionType *string `json:"exclusionType,omitempty"`
	} `json:"data"`
}