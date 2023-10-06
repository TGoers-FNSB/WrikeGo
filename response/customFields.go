package wrikeresponse

import "encoding/json"

func CustomFieldsFromJSON(data []byte) (CustomFields, error) {
	var item CustomFields
	err := json.Unmarshal(data, &item)
	return item, err
}

type CustomFields struct {
	Kind string `json:"kind"`
	Data []struct {
		Id        string        `json:"id"`
		AccountId string        `json:"accountId"`
		Title     string        `json:"title"`
		Type      string        `json:"type"`
		SharedIds []interface{} `json:"sharedIds"`
		Settings  struct {
			InheritanceType       string `json:"inheritanceType"`
			DecimalPlaces         int    `json:"decimalPlaces"`
			UseThousandsSeparator bool   `json:"useThousandsSeparator"`
			Aggregation           string `json:"aggregation"`
			ReadOnly              bool   `json:"readOnly"`
		} `json:"settings"`
	} `json:"data"`
}