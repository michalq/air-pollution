package repositories

import (
	"air-pollution/modules/core/models"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/michalq/go-airly-api-client/client"
	"github.com/michalq/go-airly-api-client/client/measurements"
	airlyModels "github.com/michalq/go-airly-api-client/models"
	"strconv"
	"time"
)

type AcquisitionRepository struct {
	client *client.AirlyAPIClient
	auth   runtime.ClientAuthInfoWriter
}

func NewAcquisitionRepository(client *client.AirlyAPIClient, auth runtime.ClientAuthInfoWriter) *AcquisitionRepository {
	return &AcquisitionRepository{client, auth}
}

func (a *AcquisitionRepository) FindAllByStationID(stationId string) ([]models.Acquisition, error) {

	indexType := "AIRLY_CAQI"
	stationIDint, _ := strconv.ParseInt(stationId, 10, 32)
	params := &measurements.InstallationMeasurementsUsingGETParams{
		IndexType:      &indexType,
		InstallationID: int32(stationIDint),
	}
	params.WithTimeout(5 * time.Second)
	resp, err := a.client.Measurements.InstallationMeasurementsUsingGET(params, a.auth)
	if err != nil {
		return nil, err
	}
	return a.translateAcquisitions(resp.Payload), nil
}

func (a *AcquisitionRepository) translateAcquisitions(
	airlyMeasurements *airlyModels.Measurements,
) []models.Acquisition {

	acqs := make([]models.Acquisition, 0)
	acqs = append(acqs, a.translateAcquisition(airlyMeasurements.Current)...)
	for _, averagedValue := range airlyMeasurements.History {
		acqs = append(acqs, a.translateAcquisition(averagedValue)...)
	}
	return acqs
}

func (a *AcquisitionRepository) translateAcquisition(
	averagedValues *airlyModels.AveragedValues,
) []models.Acquisition {

	acqs := make([]models.Acquisition, 0)
	for _, averagedValue := range averagedValues.Values {

		dateFrom, _ := time.Parse("2006-01-02 15:04:05", averagedValues.FromDateTime.String())
		dateTo, _ := time.Parse("2006-01-02 15:04:05", averagedValues.TillDateTime.String())
		acqs = append(acqs, models.Acquisition{
			Type:     models.Type(*averagedValue.Name),
			DateFrom: dateFrom,
			DateTo:   dateTo,
			Value:    fmt.Sprintf("%f", *averagedValue.Value),
		})
	}
	return acqs
}
