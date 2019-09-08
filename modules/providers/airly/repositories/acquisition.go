package repositories

import (
	"air-pollution/modules/core/models"
	"github.com/michalq/go-airly-api-client/client"
)

type AcquisitionRepository struct {
	client *client.AirlyAPIClient
}

func (a *AcquisitionRepository) FindAllByStationID(string) ([]*models.Acquisition, error) {
	return nil, nil
}
