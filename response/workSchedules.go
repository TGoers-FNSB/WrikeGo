package wrikeresponse

import (
	"encoding/json"
)

func WorkSchedulesFromJSON(data []byte) (WorkSchedules, error) {
	var item WorkSchedules
	err := json.Unmarshal(data, &item)
	return item, err
}

type WorkSchedules struct {
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
