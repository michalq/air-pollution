package repositories

import (
	"air-pollution/modules/core/models"
	giosModel "air-pollution/modules/providers/gios/models"
	"fmt"
	"github.com/michalq/go-gios-api-client/client"
	"github.com/michalq/go-gios-api-client/client/sensor"
	"github.com/michalq/go-gios-api-client/client/stations"
	"strconv"
	"time"
)

type AcquisitionRepository struct {
	client *client.GiosAPIClient
}

func NewAcquisitionRepository(client *client.GiosAPIClient) *AcquisitionRepository {
	return &AcquisitionRepository{client}
}

func (a *AcquisitionRepository) FindAllByStationID(stationID string) ([]models.Acquisition, error) {
	acqsChan := make(chan []models.Acquisition)
	errsChan := make(chan error)
	sensors, err := a.findSensorsByStationID(stationID)
	acqs := make([]models.Acquisition, 0)
	if err != nil {
		return nil, err
	}
	for _, singleSensor := range sensors {
		go func(singleSensor giosModel.Sensor) {
			intSensorID, _ := strconv.Atoi(singleSensor.ID)

			params := &sensor.SensorParams{SensorID: int64(intSensorID)}
			params.WithTimeout(5 * time.Second)
			sensorData, err := a.client.Sensor.Sensor(params)
			localAcqs := make([]models.Acquisition, 0)
			if err != nil {
				errsChan <- err
				return
			}
			for _, acqValue := range sensorData.Payload.Values {
				date, _ := time.Parse("2006-01-02 15:04:05", acqValue.Date)
				localAcqs = append(localAcqs, models.Acquisition{
					Type:     models.Type(sensorData.Payload.Key),
					DateFrom: date,
					DateTo:   date,
					Value:    fmt.Sprintf("%f", acqValue.Value),
				})
			}
			acqsChan <- localAcqs
		}(singleSensor)
	}

	defer close(acqsChan)
	for range sensors {
		select {
		case localAcqs := <-acqsChan:
			acqs = append(acqs, localAcqs...)
		case err := <-errsChan:
			fmt.Printf(err.Error())
		}
	}

	return acqs, nil
}

// TODO this must be somehow cached, this data doesn't change often.
func (a *AcquisitionRepository) findSensorsByStationID(stationID string) ([]giosModel.Sensor, error) {
	intStationID, _ := strconv.Atoi(stationID)
	params := &stations.SensorDataParams{StationID: int64(intStationID)}
	params.WithTimeout(5 * time.Second)
	rawSensors, err := a.client.Stations.SensorData(params)
	sensors := make([]giosModel.Sensor, 0)
	if err != nil {
		return nil, err
	}
	for _, rawSensor := range rawSensors.Payload {
		sensors = append(sensors, giosModel.Sensor{
			ID:        strconv.FormatInt(rawSensor.ID, 10),
			StationID: strconv.FormatInt(rawSensor.StationID, 10),
			Code:      string(rawSensor.Param.ParamCode),
		})
	}
	return sensors, nil
}
