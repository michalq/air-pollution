package repositories

import (
	"air-pollution/modules/core/models"
	"github.com/michalq/go-gios-api-client/client"
	"github.com/michalq/go-gios-api-client/client/stations"
	"strconv"
	"time"
)

type StationRepository struct {
	client *client.GiosAPIClient
}

func NewStationRepository(apiClient *client.GiosAPIClient) *StationRepository {
	return &StationRepository{apiClient}
}

func (s *StationRepository) FindAll() ([]*models.Station, error) {
	params := &stations.AllStationsParams{}
	params.WithTimeout(5 * time.Second)
	allStations, err := s.client.Stations.AllStations(params)

	if err != nil {
		return nil, err
	}

	legacyStations := make([]*models.Station, 0)
	for _, station := range allStations.Payload {

		address := &models.Address{}
		if station.City != nil {
			address.ExternalID = strconv.FormatInt(station.City.ID, 10)
			address.City = station.City.Name
			address.Country = "Poland"
			address.Street = parseNullableString(station.City.AddressStreet)
		}

		location := &models.Location{
			Latitude:  station.GegrLat,
			Longitude: station.GegrLon,
		}

		legacyStations = append(legacyStations, &models.Station{
			ExternalID: strconv.FormatInt(station.ID, 10),
			Provider:   models.Gios,
			Name:       station.StationName,
			Location:   location,
			Address:    address,
		})
	}

	return legacyStations, nil
}

func parseNullableString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}
