package repositories

import "air-pollution/daq/models"

type StationFinder interface {
	FindAll() ([]*models.Station, error)
}
