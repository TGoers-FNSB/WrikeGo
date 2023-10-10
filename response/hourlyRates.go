package wrikeresponse

import "encoding/json"

func HourlyRatesFromJSON(data []byte) (HourlyRates, error) {
	var item HourlyRates
	err := json.Unmarshal(data, &item)
	return item, err
}

type HourlyRates struct {
	Kind string `json:"kind"`
	Data []struct {
		RateSubjectId   *string `json:"rateSubjectId,omitempty"`
		RateSubjectType *string `json:"rateSubjectType,omitempty"`
		BillRate        *struct {
			RateScope  *string `json:"rateScope,omitempty"`
			RateSource *string `json:"rateSource,omitempty"`
			RateValue  *int    `json:"rateValue,omitempty"`
		} `json:"billRate,omitempty"`
		CostRate *struct {
			RateScope  *string `json:"rateScope,omitempty"`
			RateSource *string `json:"rateSource,omitempty"`
			RateValue  *int    `json:"rateValue,omitempty"`
		} `json:"costRate,omitempty"`
	} `json:"data"`
}