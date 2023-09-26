package wrikeparams

/*
	This file contains structs that are shared across multiple methods.
*/

type DateOrRange struct {
	Start string `url:"start,omitempty"`
	End   string `url:"end,omitempty"`
	Equal string `url:"equal,omitempty"`
}

type TaskDates struct {
	Type           string `url:"type,omitempty"`
	Duration       int    `url:"duration,omitempty"`
	Start          string `url:"start,omitempty"`
	Due            string `url:"due,omitempty"`
	WorkOnWeekends bool   `url:"workOnWeekends,omitempty"`
}

type Metadata struct {
	Key   string `url:"key,omitempty"`
	Value string `url:"value,omitempty"`
}

type CustomField struct {
	Id         string   `url:"id,omitempty"`
	Comparator string   `url:"comparator,omitempty"`
	Value      string   `url:"value,omitempty"`
	MinValue   string   `url:"minValue,omitempty"`
	MaxValue   string   `url:"maxValue,omitempty"`
	Values     []string `url:"values,omitempty"`
}

type EffortAllocation struct {
	AllocatedEffort           int    `url:"allocatedEffort,omitempty"`
	DailyAllocationPercentage int    `url:"dailyAllocationPercentage,omitempty"`
	Mode                      string `url:"mode,omitempty"`
	TotalEffort               int    `url:"totalEffort,omitempty"`
}

type Project struct {
	OwnerIds       []string `url:"ownerIds,omitempty"`
	Status         string   `url:"status,omitempty"`
	CustomStatusID string   `url:"customStatusId,omitempty"`
	StartDate      string   `url:"startDate,omitempty"`
	EndDate        string   `url:"endDate,omitempty"`
	ContractType   string   `url:"contractType,omitempty"`
	Budget         float64  `url:"budget,omitempty"`
	OwnersAdd      []string `url:"ownersAdd,omitempty"`
	OwnersRemove   []string `url:"ownersRemove,omitempty"`
}

type Avatar struct {
	Letters string `url:"letters,omitempty"`
	Color   string `url:"color,omitempty"`
}
