package models

type Station struct {
	ExternalID string
	Provider Provider
	Name string
	Location *Location
	Address *Address
}