package wrikeresponse

import "encoding/json"

func ContactsFromJSON(data []byte) (Contacts, error) {
	var item Contacts
	err := json.Unmarshal(data, &item)
	return item, err
}

type Contacts struct {
	Kind string `json:"kind,omitempty"`
	Data []struct {
		Id        string `json:"id,omitempty"`
		FirstName string `json:"firstName,omitempty"`
		LastName  string `json:"lastName,omitempty"`
		Type      string `json:"type,omitempty"`
		Profiles  []struct {
			AccountID string `json:"accountId,omitempty"`
			Role      string `json:"role,omitempty"`
			External  bool   `json:"external,omitempty"`
			Admin     bool   `json:"admin,omitempty"`
			Owner     bool   `json:"owner,omitempty"`
		} `json:"profiles,omitempty"`
		AvatarURL string   `json:"avatarUrl,omitempty"`
		Timezone  string   `json:"timezone,omitempty"`
		Locale    string   `json:"locale,omitempty"`
		Deleted   bool     `json:"deleted,omitempty"`
		Me        bool     `json:"me,omitempty,omitempty"`
		MemberIds []string `json:"memberIds,omitempty"`
		Metadata  []struct {
			Key   string `json:"key,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"metadata,omitempty"`
		MyTeam          bool   `json:"myTeam,omitempty"`
		Title           string `json:"title,omitempty"`
		CompanyName     string `json:"companyName,omitempty"`
		Phone           string `json:"phone,omitempty"`
		Location        string `json:"location,omitempty"`
		WorkScheduleId  string `json:"workScheduleId,omitempty"`
		CurrentBillRate struct {
			RateSource string `json:"rateSource,omitempty"`
			RateValue  string `json:"rateValue,omitempty"`
		} `json:"currentBillRate,omitempty"`
		CurrentCostRate struct {
			RateSource string `json:"rateSource,omitempty"`
			RateValue  string `json:"rateValue,omitempty"`
		} `json:"currentCostRate,omitempty"`
		JobRoleId string `json:"jobRoleId,omitempty"`
		BillRateHistory []struct {
			RateSource string `json:"rateSource,omitempty"`
			RateValue  int    `json:"rateValue,omitempty"`
			StartDate  string `json:"startDate,omitempty"`
			EndDate  string `json:"endDate,omitempty"`
		} `json:"billRateHistory,omitempty"`
		CostRateHistory []struct {
			RateSource string `json:"rateSource,omitempty"`
			RateValue  int    `json:"rateValue,omitempty"`
			StartDate  string `json:"startDate,omitempty"`
			EndDate  string `json:"endDate,omitempty"`
		} `json:"costRateHistory,omitempty"`
	} `json:"data,omitempty"`
}
