package wrikeparams

type ModifyHourlyRatesProvision struct {
	UserRates []struct {
		Contact    string   `url:"contact"`
		RateType   string   `url:"rateType"`
		RateValue  *float64 `url:"rateValue,omitempty"`
		RateSource string   `url:"rateSource"`
	} `url:"userRates"`
}
