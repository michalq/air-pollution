package repositories

import (
	"air-pollution/modules/core/models"
	giosModel "air-pollution/modules/providers/gios/models"
	"github.com/michalq/go-gios-api-client/client"
	"github.com/michalq/go-gios-api-client/client/sensor"
	"github.com/michalq/go-gios-api-client/client/stations"
	"strconv"
	"sync"
	"time"
)

type AcquisitionRepository struct {
	client *client.GiosAPIClient
}

func (a *AcquisitionRepository) FindAllByStationID(stationID string) ([]models.Acquisition, error) {
	acqsChan := make(chan []models.Acquisition)
	sensors := a.findSensorsByStationID(stationID)
	acqs := make([]models.Acquisition, 0)

	var wg sync.WaitGroup
	for _, singleSensor := range sensors {
		wg.Add(1)
		go func(singleSensor giosModel.Sensor, wg *sync.WaitGroup) {
			intSensorID, _ := strconv.Atoi(singleSensor.ID)

			params := &sensor.SensorParams{SensorID: int64(intSensorID)}
			sensorData, err := a.client.Sensor.Sensor(params)
			localAcqs := make([]models.Acquisition, 0)
			if err != nil {
				acqsChan <- localAcqs
				// TODO what to do with it?
				return
			}
			for _, acqValue := range sensorData.Payload.Values {
				date, _ := time.Parse("2006-01-02 15:04:05", acqValue.Date)
				localAcqs = append(localAcqs, models.Acquisition{
					Type:     sensorData.Payload.Key,
					DateFrom: date,
					DateTo:   date,
					Value:    acqValue.Value,
				})
			}
			acqsChan <- localAcqs
		}(singleSensor, &wg)
	}

	defer close(acqsChan)
	for range sensors {
		acqs = append(acqs, <-acqsChan...)
	}

	return acqs, nil
}

// TODO this must be somehow cached, this data doesn't change often.
func (a *AcquisitionRepository) findSensorsByStationID(stationID string) []giosModel.Sensor {
	intStationID, _ := strconv.Atoi(stationID)
	params := &stations.SensorDataParams{StationID: int64(intStationID)}
	rawSensors, err := a.client.Stations.SensorData(params)
	sensors := make([]giosModel.Sensor, 0)
	if err != nil {
		return sensors
	}
	for _, rawSensor := range rawSensors.Payload {
		sensors = append(sensors, giosModel.Sensor{
			ID:        strconv.FormatInt(rawSensor.ID, 10),
			StationID: strconv.FormatInt(rawSensor.StationID, 10),
			Code:      string(rawSensor.Param.ParamCode),
		})
	}

	return sensors
}
