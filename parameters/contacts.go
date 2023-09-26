package wrikeparams

type QueryContacts struct {
	Me          *bool        `url:"me,omitempty"` // Required in some queries
	Metadata    *Metadata    `url:"metadata,omitempty"`
	Deleted     *bool        `url:"deleted,omitempty"`
	Fields      *[]string    `url:"fields,omitempty"`
	UpdatedDate *DateOrRange `url:"updatedDate,omitempty"`
}

type ModifyContacts struct {
	Metadata        *Metadata `url:"metadata,omitempty"`
	CurrentBillRate *struct {
		RateSource *string `url:"rateSource,omitempty"`
		RateValue  *string `url:"rateValue,omitempty"`
	} `url:"currentBillRate,omitempty"`
	CurrentCostRate *struct {
		RateSource *string `url:"rateSource,omitempty"`
		RateValue  *string `url:"rateValue,omitempty"`
	} `url:"currentCostRate,omitempty"`
	JobRoleId *string   `url:"jobRoleId,omitempty"`
	Fields    *[]string `url:"fields,omitempty"`
}
