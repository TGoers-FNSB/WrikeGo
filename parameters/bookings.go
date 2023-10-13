package wrikeparams

type QueryBookings struct {
	StartDate                 string   `url:"startDate,omitempty"`
	FinishDate                string   `url:"finishDate,omitempty"`
	ResponsibleIds            []string `url:"responsibleIds,omitempty,slice"`
	ResponsiblePlaceholderIds []string `url:"respopnsiblePlaceholderIds,omitempty,slice"`
	ShowDescendants           string   `url:"showDescendants,omitempty"`
}

type CreateBookings struct {
	BookingDates             BookingDates     `url:"bookingDates,struct"`
	ResponsibleId            string           `url:"responsibleId,omitempty"`
	ResponsiblePlaceholderId string           `url:"respopnsiblePlaceholderId,omitempty"`
	EffortAllocation         EffortAllocation `url:"effortAllocation,omitempty,struct"`
}

type ModifyBookings struct {
	BookingDates             BookingDates     `url:"bookingDates,omitempty,struct"`
	ResponsibleId            string           `url:"responsibleId,omitempty"`
	ResponsiblePlaceholderId string           `url:"respopnsiblePlaceholderId,omitempty"`
	EffortAllocation         EffortAllocation `url:"effortAllocation,omitempty,struct"`
}
