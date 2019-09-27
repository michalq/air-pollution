package repositories

import "air-pollution/daq/core/models"

type AcquisitionFinder interface {
	FindAllByStationID(string) ([]*models.Acquisition, error)
}

type AcquisitionPersister interface {
	Persist(*models.Acquisition) error
}
