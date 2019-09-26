package repositories

import "air-pollution/modules/core/models"

type AcquisitionFinder interface {
	FindAllByStationID(string) ([]*models.Acquisition, error)
}

type AcquisitionSaver interface {
	Save(*models.Acquisition) error
}
