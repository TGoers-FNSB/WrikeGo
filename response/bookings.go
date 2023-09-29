package wrikeresponse

import (
	"encoding/json"
)

func BookingsFromJSON(data []byte) (Bookings, error) {
	var item Bookings
	err := json.Unmarshal(data, &item)
	return item, err
}

type Bookings struct {
	Kind string `json:"kind"`
	Data []struct {
		ID            string `json:"id"`
		FolderID      string `json:"folderId"`
		ResponsibleID string `json:"responsibleId"`
		BookingDates  struct {
			Duration       int    `json:"duration"`
			StartDate      string `json:"startDate"`
			FinishDate     string `json:"finishDate"`
			WorkOnWeekends bool   `json:"workOnWeekends"`
		} `json:"bookingDates"`
		EffortAllocation struct {
			ResponsibleAllocation []struct {
				UserID          string        `json:"userId"`
				DailyAllocation []interface{} `json:"dailyAllocation"`
			} `json:"responsibleAllocation"`
			Mode        string `json:"mode"`
			TotalEffort int    `json:"totalEffort"`
		} `json:"effortAllocation"`
	} `json:"data"`
}