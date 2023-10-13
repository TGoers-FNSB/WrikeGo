package wrikeparams

type QueryContacts struct {
	Me          bool        `url:"me,omitempty"` //! Required in some queries
	Metadata    Metadata    `url:"metadata,omitempty,struct"`
	Deleted     bool        `url:"deleted,omitempty"`
	Fields      []string    `url:"fields,omitempty,slice"`
	UpdatedDate DateOrRange `url:"updatedDate,omitempty,struct"`
}

type ModifyContacts struct {
	Metadata        []Metadata `url:"metadata,omitempty,slice+struct"`
	CurrentBillRate struct {
		RateSource string `url:"rateSource"`
		RateValue  string `url:"rateValue,omitempty"`
	} `url:"currentBillRate,omitempty,struct"`
	CurrentCostRate struct {
		RateSource string `url:"rateSource"`
		RateValue  string `url:"rateValue,omitempty"`
	} `url:"currentCostRate,omitempty,struct"`
	JobRoleId string   `url:"jobRoleId,omitempty"`
	Fields    []string `url:"fields,omitempty,slice"`
}
