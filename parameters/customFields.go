package wrikeparams

type CreateCustomFields struct {
	Title	string	`url:"title"`
	Type	string	`url:"type"`
	SpaceId	*string	`url:"spaceId,omitempty"`
	Shareds	*[]string	`url:"shareds,omitempty"`
	Settings  *struct {
		InheritanceType       *string   `url:"inheritanceType,omitempty"`
		DecimalPlaces         *int      `url:"decimalPlaces,omitempty"`
		UseThousandsSeparator *bool     `url:"useThousandsSeparator,omitempty"`
		Currency              *string   `url:"currency,omitempty"`
		Aggregation           *string   `url:"aggregation,omitempty"`
		Values                *[]string `url:"values,omitempty"`
		Options               *[]string `url:"options,omitempty"`
		OptionColorsEnabled   *bool     `url:"optionColorsEnabled,omitempty"`
		AllowedOtherValues    *bool     `url:"allowedOtherValues,omitempty"`
		Contacts              *[]string `url:"contacts,omitempty"`
	} `url:"settings,omitempty"`
}

type ModifyCustomFields struct {
	Title	*string	`url:"title,omitempty"`
	Type	*string	`url:"type,omitempty"`
	ChangeScope	*string	`url:"changeScope,omitempty"`
	SpaceId	*string	`url:"spaceId,omitempty"`
	AddShareds	*[]string	`url:"addShareds,omitempty"`
	RemoveShareds	*[]string	`url:"removeShareds,omitempty"`
	Settings  *struct {
		InheritanceType       *string   `url:"inheritanceType,omitempty"`
		DecimalPlaces         *int      `url:"decimalPlaces,omitempty"`
		UseThousandsSeparator *bool     `url:"useThousandsSeparator,omitempty"`
		Currency              *string   `url:"currency,omitempty"`
		Aggregation           *string   `url:"aggregation,omitempty"`
		Values                *[]string `url:"values,omitempty"`
		Options               *[]string `url:"options,omitempty"`
		OptionColorsEnabled   *bool     `url:"optionColorsEnabled,omitempty"`
		AllowedOtherValues    *bool     `url:"allowedOtherValues,omitempty"`
		Contacts              *[]string `url:"contacts,omitempty"`
	} `url:"settings,omitempty"`
}