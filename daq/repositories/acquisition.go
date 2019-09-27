package repositories

import "air-pollution/daq/models"

type AcquisitionFinder interface {
	FindAllByStationID(string) ([]*models.Acquisition, error)
}

type AcquisitionPersister interface {
	Persist(*models.Acquisition) error
}
