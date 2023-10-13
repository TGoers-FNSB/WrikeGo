package wrikeparams

type ModifyHourlyRates struct {
	Rates []struct {
		RateSubjectId string  `url:"rateSubjectId"`
		RateType      string  `url:"rateType"`
		RateValue     float64 `url:"rateValue,omitempty"`
		RateSource    string  `url:"rateSource,omitempty"`
	} `url:"rates,slice+struct"`
}

type DeleteTeamMembers struct {
	RateSubjects []string `url:"rateSubjects,slice"`
}
