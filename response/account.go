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
		Id             string   `json:"Id"`
		Name           string   `json:"name"`
		DateFormat     string   `json:"dateFormat"`
		FirstDayOfWeek string   `json:"firstDayOfWeek"`
		WorkDays       []string `json:"workDays"`
		RootFolderId   string   `json:"rootFolderId"`
		RecycleBinId   string   `json:"recycleBinId"`
		CreatedDate    string   `json:"createdDate"`
		Subscription   struct {
			Type      string `json:"type"`
			Suspended bool   `json:"suspended"`
			Paid      bool   `json:"paid"`
			UserLimit int    `json:"userLimit"`
		} `json:"subscription"`
		Metadata []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"metadata"`
		CustomFields []struct {
			Id        string   `json:"Id"`
			AccountId string   `json:"accountId"`
			Title     string   `json:"title"`
			Type      string   `json:"type"`
			SpaceId   string   `json:"spaceId"`
			SharedIds []string `json:"sharedIds"`
			Settings  struct {
				InheritanceType       string   `json:"inheritanceType"`
				DecimalPlaces         int      `json:"decimalPlaces"`
				UseThousandsSeparator bool     `json:"useThousandsSeparator"`
				Currency              string   `json:"currency"`
				Aggregation           string   `json:"aggregation"`
				Values                []string `json:"values"`
				Options               []string `json:"options"`
				OptionColorsEnabled   bool     `json:"optionColorsEnabled"`
				AllowedOtherValues    bool     `json:"allowedOtherValues"`
				Contacts              []string `json:"contacts"`
				ReadOnly              bool     `json:"readOnly"`
				AllowTime             bool     `json:"allowTime"`
				Timezone              string   `json:"timezone"`
			} `json:"settings"`
		} `json:"customFields"`
		JoinedDate string `json:"joinedDate"`
	} `json:"data"`
}
