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
	Type           string `url:"type"`
	Duration       int    `url:"duration,omitempty"`
	Start          string `url:"start,omitempty"`
	Due            string `url:"due,omitempty"`
	WorkOnWeekends bool   `url:"workOnWeekends,omitempty"`
}

type Metadata struct {
	Key   string `url:"key"`
	Value string `url:"value"`
}

type CustomField struct {
	Id         string   `url:"id"`
	Value      string   `url:"value"`
	Comparator string   `url:"comparator,omitempty"`
	MinValue   string   `url:"minValue,omitempty"`
	MaxValue   string   `url:"maxValue,omitempty"`
	Values     []string `url:"values,omitempty,slice"`
}

type Project struct {
	OwnerIds       []string `url:"ownerIds,omitempty,slice"`
	Status         string   `url:"status,omitempty"`
	CustomStatusId string   `url:"customStatusId,omitempty"`
	StartDate      string   `url:"startDate,omitempty"`
	EndDate        string   `url:"endDate,omitempty"`
	ContractType   string   `url:"contractType,omitempty"`
	Budget         float64  `url:"budget,omitempty"`
	OwnersAdd      []string `url:"ownersAdd,omitempty,slice"`
	OwnersRemove   []string `url:"ownersRemove,omitempty,slice"`
}

type Avatar struct {
	Letters string `url:"letters"`
	Color   string `url:"color"`
}

type BookingDates struct {
	Duration       int    `url:"duration,omitempty"`
	StartDate      string `url:"startDate,omitempty"`
	FinishDate     string `url:"finishDate,omitempty"`
	WorkOnWeekends bool   `url:"workOnWeekends,omitempty"`
}

type EffortAllocation struct {
	AllocatedEffort           int                     `url:"allocatedEffort,omitempty"`
	DailyAllocationPercentage int                     `url:"dailyAllocationPercentage,omitempty"`
	Mode                      string                  `url:"mode"`
	TotalEffort               int                     `url:"totalEffort,omitempty"`
	ResponsibleAllocation     []ResponsibleAllocation `url:"responsibleAllocation,omitempty,slice+struct"`
}

type ResponsibleAllocation struct {
	UserId          string   `url:"userId,omitempty"`
	PlaceholderId   string   `url:"placeholderId,omitempty"`
	DailyAllocation []string `url:"dailyAllocation,slice"` //? Data Type
}

type Settings struct {
	InheritanceType       string   `url:"inheritanceType,omitempty"`
	DecimalPlaces         int      `url:"decimalPlaces,omitempty"`
	UseThousandsSeparator bool     `url:"useThousandsSeparator,omitempty"`
	Currency              string   `url:"currency,omitempty"`
	Aggregation           string   `url:"aggregation,omitempty"`
	Values                []string `url:"values,omitempty,slice"`
	Options               []string `url:"options,omitempty,slice"`
	OptionColorsEnabled   bool     `url:"optionColorsEnabled,omitempty"`
	AllowOtherValues      bool     `url:"allowOtherValues,omitempty"`
	Contacts              []string `url:"contacts,omitempty,slice"`
}

type Members struct {
	Id           string `url:"id"`
	AccessRoleId string `url:"accessRoleId"`
	IsManager    bool   `url:"isManager"`
}
