package wrikeresponse

import "encoding/json"

func GroupsFromJSON(data []byte) (Groups, error) {
	var item Groups
	err := json.Unmarshal(data, &item)
	return item, err
}

type Groups struct {
	Kind string `json:"kind,omitempty"`
	Data []struct {
		Id        *string   `json:"id,omitempty"`
		AccountId *string   `json:"accountId,omitempty"`
		Title     *string   `json:"title,omitempty"`
		MemberIds *[]string `json:"memberIds,omitempty"`
		ChildIds  *[]string `json:"childIds,omitempty"`
		ParentIds *[]string `json:"parentIds,omitempty"`
		AvatarURL *string   `json:"avatarUrl,omitempty"`
		MyTeam    *bool     `json:"myTeam,omitempty"`
		Metadata  *[]struct {
			Key   *string `json:"key,omitempty"`
			Value *string `json:"value,omitempty"`
		} `json:"metadata,omitempty"`
	} `json:"data,omitempty"`
}
