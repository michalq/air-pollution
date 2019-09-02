package models

type Station struct {

	// External identifier giver by provider
	ExternalID string
	Provider   Provider
	Name       string
	Location   *Location
	Address    *Address
}
