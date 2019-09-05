package repositories

import (
	"air-pollution/modules/core/models"
	"github.com/michalq/go-gios-api-client/client"
	"github.com/michalq/go-gios-api-client/client/sensor"
	"strconv"
)

type AcquisitionRepository struct {
	client *client.GiosAPIClient
}

func (s *AcquisitionRepository) FindAllByStationID(stationID string) ([]*models.Acquisition, error) {
	intStationID, _ := strconv.Atoi(stationID)
	params := &sensor.SensorParams{SensorID: int64(intStationID)}
	_, err := s.client.Sensor.Sensor(params)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
