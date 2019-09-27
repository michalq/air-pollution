package repositories

import "air-pollution/daq/core/models"

type StationFinder interface {
	FindAll() ([]*models.Station, error)
}
