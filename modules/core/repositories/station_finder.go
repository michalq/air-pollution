package repositories

import "air-pollution/modules/core/models"

type StationFinder interface {
	FindAll() ([]*models.Station, error)
}
