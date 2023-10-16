package wrikeresponse

import "encoding/json"

func UsersFromJSON(data []byte) (Users, error) {
	var item Users
	err := json.Unmarshal(data, &item)
	return item, err
}

type Users struct {
	Kind string `json:"kind"`
	Data []struct {
		Id        *string `json:"id,omitempty"`
		FirstName *string `json:"firstName,omitempty"`
		LastName  *string `json:"lastName,omitempty"`
		Type      *string `json:"type,omitempty"`
		Profiles  *[]struct {
			AccountId *string `json:"accountId,omitempty"`
			Email     *string `json:"email,omitempty"`
			Role      *string `json:"role,omitempty"`
			External  *bool   `json:"external,omitempty"`
			Admin     *bool   `json:"admin,omitempty"`
			Owner     *bool   `json:"owner,omitempty"`
		} `json:"profiles,omitempty"`
		AvatarURL *string `json:"avatarUrl,omitempty"`
		Timezone  *string `json:"timezone,omitempty"`
		Locale    *string `json:"locale,omitempty"`
		Deleted   *bool   `json:"deleted,omitempty"`
		Me        *bool   `json:"me,omitempty"`
		MemberIds *[]string `json:"memberIds,omitempty"`
		Metadata  *[]struct {
			Key   *string `json:"key,omitempty"`
			Value *string `json:"value,omitempty"`
		} `json:"metadata,omitempty"`
		MyTeam          *bool   `json:"myTeam,omitempty"`
		Title           *string `json:"title,omitempty"`
		CompanyName     *string `json:"companyName,omitempty"`
		Phone           *string `json:"phone,omitempty"`
		Location        *string `json:"location,omitempty"`
		WorkScheduleId  *string `json:"workScheduleId,omitempty"`
	} `json:"data"`
	// Errors:
	ErrorDescription *string `json:"errorDescription"`
	Error            *string `json:"error"`
}