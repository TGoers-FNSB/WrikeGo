package wrikeresponse

import (
	"encoding/json"
)

func CopyWorkSchedulesFromJSON(data []byte) (CopyWorkSchedules, error) {
	var item CopyWorkSchedules
	err := json.Unmarshal(data, &item)
	return item, err
}

type CopyWorkSchedules struct {
	Kind string `json:"kind"`
	Data []struct {
		Id           *string `json:"id,omitempty"`
		ScheduleType *string `json:"scheduleType,omitempty"`
		Title        *string `json:"title,omitempty"`
		Workweek     *[]struct {
			WorkDays        *[]string `json:"workDays,omitempty"`
			CapacityMinutes *int      `json:"capacityMinutes,omitempty"`
		} `json:"workweek,omitempty"`
		UserIds *[]string `json:"userIds,omitempty"`
	} `json:"data"`
}
