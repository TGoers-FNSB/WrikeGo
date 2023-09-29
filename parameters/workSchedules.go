package wrikeparams

type QueryWorkSchedules struct {
	Fields *[]string `url:"fields,omitempty"`
}

type CreateWorkSchedules struct {
	Title    string `url:"title"`
	Workweek []struct {
		DayOfWeek string `url:"dayOfWeek"`
		IsWorkDay bool   `url:"isWorkDay"`
	} `url:"workweek"`
	AddUsers        *[]string `url:"addUsers,omitempty"`
	CapacityMinutes *int      `url:"capacityMinutes,omitempty"`
	Fields          *[]string `url:"fields,omitempty"`
}

type ModifyWorkSchedules struct {
	Title    *string `url:"title,omitempty"`
	Workweek *[]struct {
		DayOfWeek string `url:"dayOfWeek"`
		IsWorkDay bool   `url:"isWorkDay"`
	} `url:"workweek,omitempty"`
	AddUsers        *[]string `url:"addUsers,omitempty"`
	RemoveUsers     *[]string `url:"removeUsers,omitempty"`
	CapacityMinutes *int      `url:"capacityMinutes,omitempty"`
	Fields          *[]string `url:"fields,omitempty"`
}
