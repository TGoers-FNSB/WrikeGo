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
		Id            *string `json:"id,omitempty"`
		FolderId      *string `json:"folderId,omitempty"`
		ResponsibleId *string `json:"responsibleId,omitempty"`
		BookingDates  *struct {
			Duration       *int    `json:"duration,omitempty"`
			StartDate      *string `json:"startDate,omitempty"`
			FinishDate     *string `json:"finishDate,omitempty"`
			WorkOnWeekends *bool   `json:"workOnWeekends,omitempty"`
		} `json:"bookingDates,omitempty"`
		EffortAllocation *struct {
			ResponsibleAllocation *[]struct {
				UserId          *string        `json:"userId,omitempty"`
				DailyAllocation *[]interface{} `json:"dailyAllocation,omitempty"` //? Unknown data type
			} `json:"responsibleAllocation,omitempty"`
			Mode        *string `json:"mode,omitempty"`
			TotalEffort *int    `json:"totalEffort,omitempty"`
		} `json:"effortAllocation,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}