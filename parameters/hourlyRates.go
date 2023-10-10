package wrikeparams

type ModifyHourlyRates struct {
	Rates []struct {
		RateSubjectId string   `url:"rateSubjectId"`
		RateType      string   `url:"rateType"`
		RateValue     *float64 `url:"rateValue,omitempty"`
		RateSource    *string  `url:"rateSource,omitempty"`
	}
}

type DeleteTeamMembers struct {
	RateSubjects []string `url:"rateSubjects"`
}
