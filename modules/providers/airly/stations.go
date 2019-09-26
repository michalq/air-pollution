package airly

import (
	"air-pollution/modules/core/models"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/michalq/go-airly-api-client/client"
	"github.com/michalq/go-airly-api-client/client/installations"
	"strconv"
	"time"
)

type StationRepository struct {
	client *client.AirlyAPIClient
	auth   runtime.ClientAuthInfoWriter
}

func NewStationRepository(client *client.AirlyAPIClient, auth runtime.ClientAuthInfoWriter) *StationRepository {
	return &StationRepository{client, auth}
}

func (s *StationRepository) FindAll() ([]*models.Station, error) {
	var maxDist float64
	var maxResults int32
	maxDist = -1
	maxResults = -1
	params := &installations.NearestUsingGETParams{
		Lat:           0,
		Lng:           0,
		MaxDistanceKM: &maxDist,
		MaxResults:    &maxResults,
	}
	params.WithTimeout(5 * time.Second)
	inst, err := s.client.Installations.NearestUsingGET(params, s.auth)
	if err != nil {
		return nil, err
	}

	fmt.Println(*inst.GetPayload()[0].Address.Country)
	legacyStations := make([]*models.Station, 0)
	for _, installation := range inst.GetPayload() {

		address := &models.Address{}
		if installation.Address != nil {
			address.ExternalID = ""
			address.City = parseNullableString(installation.Address.City)
			address.Country = models.Country(parseNullableString(installation.Address.Country))
			address.Street = parseNullableString(installation.Address.Street)
		}

		location := &models.Location{
			Latitude:  strconv.FormatFloat(installation.Location.Latitude, 'f', 10, 64),
			Longitude: strconv.FormatFloat(installation.Location.Longitude, 'f', 10, 64),
		}

		legacyStations = append(legacyStations, &models.Station{
			ExternalID: strconv.FormatInt(int64(installation.ID), 10),
			Provider:   models.Airly,
			Name: fmt.Sprintf(
				"%s, %s, %s %s",
				parseNullableString(installation.Address.Country),
				parseNullableString(installation.Address.City),
				parseNullableString(installation.Address.DisplayAddress1),
				parseNullableString(installation.Address.DisplayAddress2),
			),
			Location: location,
			Address:  address,
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
