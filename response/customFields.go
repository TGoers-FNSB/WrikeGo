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
		Id        *string        `json:"id,omitempty"`
		AccountId *string        `json:"accountId,omitempty"`
		Title     *string        `json:"title,omitempty"`
		Type      *string        `json:"type,omitempty"`
		SharedIds *[]interface{} `json:"sharedIds,omitempty"`
		Settings  *struct {
			InheritanceType       *string `json:"inheritanceType,omitempty"`
			DecimalPlaces         *int    `json:"decimalPlaces,omitempty"`
			UseThousandsSeparator *bool   `json:"useThousandsSeparator,omitempty"`
			Aggregation           *string `json:"aggregation,omitempty"`
			ReadOnly              *bool   `json:"readOnly,omitempty"`
		} `json:"settings,omitempty"`
	} `json:"data"`
}