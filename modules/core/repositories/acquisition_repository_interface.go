package repositories

import "air-pollution/modules/core/models"

type AcquisitionRepositoryInterface interface {
	FindAllByStationID(string) ([]*models.Acquisition, error)
}
