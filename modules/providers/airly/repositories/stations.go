package repositories

import (
	"air-pollution/modules/core/models"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/michalq/go-airly-api-client/client"
	"github.com/michalq/go-airly-api-client/client/installations"
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
		Lat:           50.0646501,
		Lng:           19.9449799,
		MaxDistanceKM: &maxDist,
		MaxResults:    &maxResults,
	}
	params.WithTimeout(5 * time.Second)
	inst, err := s.client.Installations.NearestUsingGET(params, s.auth)
	if err != nil {
		return nil, err
	}

	fmt.Println(inst)
	return nil, nil
}
