// Package listbooks provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package listbooks

import (
	externalRef0 "github.com/srijilv/go-api-template.git/pkg/interfaces/openapi/common"
)

// BookAttributes defines model for BookAttributes.
type BookAttributes struct {
	Author        string `json:"author"`
	DateOfPublish string `json:"dateOfPublish"`
	Pages         int16  `json:"pages"`
	Price         string `json:"price"`
	Publisher     string `json:"publisher"`
}

// Facility defines model for Facility.
type Facility struct {
	Billing ServiceLocation `json:"billing"`
	Service ServiceLocation `json:"service"`
}

// ListBooksResponse defines model for ListBooksResponse.
type ListBooksResponse struct {
	Info    externalRef0.Information   `json:"info"`
	Payload []ListBooksResponsePayload `json:"payload"`
}

// ListBooksResponsePayload defines model for ListBooksResponsePayload.
type ListBooksResponsePayload struct {
	Attributes BookAttributes `json:"attributes"`
	Id         int32          `json:"id"`
	LongDesc   string         `json:"longDesc"`
	Name       string         `json:"name"`
	ShortDesc  string         `json:"shortDesc"`
}

// ServiceLocation defines model for ServiceLocation.
type ServiceLocation struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	PostalCode string `json:"postalCode"`
	State      string `json:"state"`
}
