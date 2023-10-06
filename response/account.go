package wrikeresponse

import "encoding/json"

func AccountFromJSON(data []byte) (Account, error) {
	var item Account
	err := json.Unmarshal(data, &item)
	return item, err
}

type Account struct {
	Kind string `json:"kind"`
	Data []struct {
		Id             *string   `json:"Id,omitempty"`
		Name           *string   `json:"name,omitempty"`
		DateFormat     *string   `json:"dateFormat,omitempty"`
		FirstDayOfWeek *string   `json:"firstDayOfWeek,omitempty"`
		WorkDays       *[]string `json:"workDays,omitempty"`
		RootFolderId   *string   `json:"rootFolderId,omitempty"`
		RecycleBinId   *string   `json:"recycleBinId,omitempty"`
		CreatedDate    *string   `json:"createdDate,omitempty"`
		Subscription   *struct {
			Type      *string `json:"type,omitempty"`
			Suspended *bool   `json:"suspended,omitempty"`
			Paid      *bool   `json:"paid,omitempty"`
			UserLimit *int    `json:"userLimit,omitempty"`
		} `json:"subscription,omitempty"`
		Metadata *[]struct {
			Key   *string `json:"key,omitempty"`
			Value *string `json:"value,omitempty"`
		} `json:"metadata,omitempty"`
		CustomFields *[]struct {
			Id        *string   `json:"Id,omitempty"`
			AccountId *string   `json:"accountId,omitempty"`
			Title     *string   `json:"title,omitempty"`
			Type      *string   `json:"type,omitempty"`
			SpaceId   *string   `json:"spaceId,omitempty"`
			SharedIds *[]string `json:"sharedIds,omitempty"`
			Settings  *struct {
				InheritanceType       *string   `json:"inheritanceType,omitempty"`
				DecimalPlaces         *int      `json:"decimalPlaces,omitempty"`
				UseThousandsSeparator *bool     `json:"useThousandsSeparator,omitempty"`
				Currency              *string   `json:"currency,omitempty"`
				Aggregation           *string   `json:"aggregation,omitempty"`
				Values                *[]string `json:"values,omitempty"`
				Options               *[]string `json:"options,omitempty"`
				OptionColorsEnabled   *bool     `json:"optionColorsEnabled,omitempty"`
				AllowedOtherValues    *bool     `json:"allowedOtherValues,omitempty"`
				Contacts              *[]string `json:"contacts,omitempty"`
				ReadOnly              *bool     `json:"readOnly,omitempty"`
				AllowTime             *bool     `json:"allowTime,omitempty"`
				Timezone              *string   `json:"timezone,omitempty"`
			} `json:"settings,omitempty"`
		} `json:"customFields,omitempty"`
		JoinedDate *string `json:"joinedDate,omitempty"`
	} `json:"data"`
}
