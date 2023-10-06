package wrikeparams

type QueryBookings struct {
	StartDate                 *string   `url:"startDate,omitempty"`
	FinishDate                *string   `url:"finishDate,omitempty"`
	ResponsibleIds            *[]string `url:"responsibleIds,omitempty"`
	ResponsiblePlaceholderIds *[]string `url:"respopnsiblePlaceholderIds,omitempty"`
	ShowDescendants           *string   `url:"showDescendants,omitempty"`
}

type CreateBookings struct {
	BookingDates struct {
		Duration       *int    `url:"duration,omitempty"`
		StartDate      *string `url:"startDate,omitempty"`
		FinishDate     *string `url:"finishDate,omitempty"`
		WorkOnWeekends *bool   `url:"workOnWeekends,omitempty"`
	} `url:"bookingDates"`
	ResponsibleId            *string `url:"responsibleId,omitempty"`
	ResponsiblePlaceholderId *string `url:"respopnsiblePlaceholderId,omitempty"`
	EffortAllocation         *struct {
		ResponsibleAllocation *[]struct {
			UserId          *string  `url:"userId,omitempty"`
			PlaceholderId   *string  `url:"placeholderId,omitempty"`
			DailyAllocation []string `url:"dailyAllocation"` //? Data Type
		} `url:"responsibleAllocation,omitempty"`
		Mode        string `url:"mode"`
		TotalEffort *int   `url:"totalEffort,omitempty"`
	} `url:"effortAllocation,omitempty"`
}

type ModifyBookings *struct {
	BookingDates struct {
		Duration       *int    `url:"duration,omitempty"`
		StartDate      *string `url:"startDate,omitempty"`
		FinishDate     *string `url:"finishDate,omitempty"`
		WorkOnWeekends *bool   `url:"workOnWeekends,omitempty"`
	} `url:"bookingDates,omitempty"`
	ResponsibleId            *string `url:"responsibleId,omitempty"`
	ResponsiblePlaceholderId *string `url:"respopnsiblePlaceholderId,omitempty"`
	EffortAllocation         *struct {
		ResponsibleAllocation *[]struct {
			UserId          *string   `url:"userId,omitempty"`
			PlaceholderId   *string   `url:"placeholderId,omitempty"`
			DailyAllocation *[]string `url:"dailyAllocation"` //? Data Type
		} `url:"responsibleAllocation,omitempty"`
		Mode        string `url:"mode"`
		TotalEffort *int   `url:"totalEffort,omitempty"`
	} `url:"effortAllocation,omitempty"`
}
