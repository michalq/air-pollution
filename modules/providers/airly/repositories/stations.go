package repositories

import (
	"air-pollution/modules/core/models"
	"github.com/michalq/go-airly-api-client/client"
)

type StationRepository struct {
	client *client.AirlyAPIClient
}

func NewStationRepository(client *client.AirlyAPIClient) *StationRepository {
	return &StationRepository{client}
}

func (s *StationRepository) FindAll() ([]*models.Station, error) {
	return s.client.Installations.ByIDUsingGET(), nil
}
