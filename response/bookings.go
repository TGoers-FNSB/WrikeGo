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
		Id            string `json:"id"`
		FolderId      string `json:"folderId"`
		ResponsibleId string `json:"responsibleId"`
		BookingDates  struct {
			Duration       int    `json:"duration"`
			StartDate      string `json:"startDate"`
			FinishDate     string `json:"finishDate"`
			WorkOnWeekends bool   `json:"workOnWeekends"`
		} `json:"bookingDates"`
		EffortAllocation struct {
			ResponsibleAllocation []struct {
				UserId          string        `json:"userId"`
				DailyAllocation []interface{} `json:"dailyAllocation"` //? Unknown data type
			} `json:"responsibleAllocation"`
			Mode        string `json:"mode"`
			TotalEffort int    `json:"totalEffort"`
		} `json:"effortAllocation"`
	} `json:"data"`
}