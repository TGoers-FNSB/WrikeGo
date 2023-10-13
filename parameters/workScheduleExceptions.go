package wrikeparams

type QueryWorkScheduleExceptions struct {
	DateRange DateOrRange `url:"dateRange,omitempty,struct"`
}

type CreateWorkScheduleExceptions struct {
	FromDate      string `url:"fromDate"`
	ToDate        string `url:"toDate"`
	ExclusionType string `url:"exclusionType"`
}

type ModifyWorkScheduleExceptions struct {
	FromDate      string `url:"fromDate,omitempty"`
	ToDate        string `url:"toDate,omitempty"`
	ExclusionType string `url:"exclusionType,omitempty"`
}
