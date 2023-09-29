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
		Id            string `json:"id"`
		UserId        string `json:"userId"`
		FromDate      string `json:"fromDate"`
		ToDate        string `json:"toDate"`
		IsWorkDays    bool   `json:"isWorkDays"`
		ExclusionType string `json:"exclusionType"`
	} `json:"data"`
}