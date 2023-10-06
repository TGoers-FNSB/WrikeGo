package wrikeparams

type QueryUserScheduleExceptions struct {
	DateRange *DateOrRange `url:"dateRange,omitempty"`
	UserIds   *[]string    `url:"userIds,omitempty"`
}

type CreateUserScheduleExceptions struct {
	UserId        interface{} `url:"userId"` //? Unknown datatype
	FromDate      string      `url:"fromDate"`
	ToDate        string      `url:"toDate"`
	ExclusionType string      `url:"exclusionType"`
}

type ModifyUserScheduleExceptions struct {
	FromDate      *string `url:"fromDate,omitempty"`
	ToDate        *string `url:"toDate,omitempty"`
	ExclusionType *string `url:"exclusionType,omitempty"`
}
