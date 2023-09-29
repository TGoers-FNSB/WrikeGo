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
		Id           string `json:"id"`
		ScheduleType string `json:"scheduleType"`
		Title        string `json:"title"`
		Workweek     []struct {
			WorkDays        []string `json:"workDays"`
			CapacityMinutes int      `json:"capacityMinutes"`
		} `json:"workweek"`
		UserIds []string `json:"userIds"`
	} `json:"data"`
}
