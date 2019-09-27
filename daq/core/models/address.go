package models

type Address struct {

	// Provider custom identifier
	ExternalID string
	Country    Country
	City       string
	// First line of street
	Street string
	// Second line of street, for example house number
	Street2 string
}
