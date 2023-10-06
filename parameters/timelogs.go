package wrikeparams

type QueryTimelogs struct {
	CreatedDate       *DateOrRange `url:"createdDate,omitempty"`
	UpdatedDate       *DateOrRange `url:"updatedDate,omitempty"`
	TrackedDate       *DateOrRange `url:"trackedDate,omitempty"`
	Me                *bool        `url:"me,omitempty"`
	Descendants       *bool        `url:"descendants,omitempty"`
	PlainText         *bool        `url:"plainText,omitempty"`
	TimelogCategories *[]string    `url:"timelogCategories,omitempty"`
	BillingTypes      *[]string    `url:"billingTypes,omitempty"`
	Fields            *[]string    `url:"fields,omitempty"`
}

type CreateTimelogs struct {
	Comment     string    `url:"comment"`
	Hours       float64   `url:"hours"`
	TrackedDate string    `url:"trackedDate"`
	PlainText   *bool     `url:"plainText,omitempty"`
	CategoryId  *string   `url:"categoryId,omitempty"`
	OnBehalfOf  *string   `url:"onBehalfOf,omitempty"`
	Fields      *[]string `url:"fields,omitempty"`
}

type ModifyTimelogs struct {
	Comment     *string   `url:"comment,omitempty"`
	Hours       *float64  `url:"hours,omitempty"`
	TrackedDate *string   `url:"trackedDate,omitempty"`
	PlainText   *bool     `url:"plainText,omitempty"`
	CategoryId  *string   `url:"categoryId,omitempty"`
	Fields      *[]string `url:"fields,omitempty"`
}
