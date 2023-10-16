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
		Contact    *string  `json:"contact,omitempty"`
		RateType   *string  `json:"rateType,omitempty"`
		RateValue  *float64 `json:"rateValue,omitempty"`
		RateSource *string  `json:"rateSource,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}