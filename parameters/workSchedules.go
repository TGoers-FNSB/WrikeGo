package wrikeparams

type QueryWorkSchedules struct {
	Fields []string `url:"fields,omitempty,slice"`
}

type CreateWorkSchedules struct {
	Title    string `url:"title"`
	Workweek []struct {
		DayOfWeek string `url:"dayOfWeek"`
		IsWorkDay bool   `url:"isWorkDay"`
	} `url:"workweek,slice+struct"`
	AddUsers        []string `url:"addUsers,omitempty,slice"`
	CapacityMinutes int      `url:"capacityMinutes,omitempty"`
	Fields          []string `url:"fields,omitempty,slice"`
}

type ModifyWorkSchedules struct {
	Title    string `url:"title,omitempty"`
	Workweek []struct {
		DayOfWeek string `url:"dayOfWeek"`
		IsWorkDay bool   `url:"isWorkDay"`
	} `url:"workweek,omitempty,slice+struct"`
	AddUsers        []string `url:"addUsers,omitempty,slice"`
	RemoveUsers     []string `url:"removeUsers,omitempty,slice"`
	CapacityMinutes int      `url:"capacityMinutes,omitempty"`
	Fields          []string `url:"fields,omitempty,slice"`
}
