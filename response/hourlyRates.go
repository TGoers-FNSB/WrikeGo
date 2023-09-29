package wrikeresponse

import (
	"encoding/json"
)

func HourlyRatesProvisionFromJSON(data []byte) (HourlyRatesProvision, error) {
	var item HourlyRatesProvision
	err := json.Unmarshal(data, &item)
	return item, err
}

type HourlyRatesProvision struct {
	Kind string `json:"kind"`
	Data []struct {
		Contact    string  `json:"contact"`
		RateType   string  `json:"rateType"`
		RateValue  float64 `json:"rateValue"`
		RateSource string  `json:"rateSource"`
	} `json:"data"`
}