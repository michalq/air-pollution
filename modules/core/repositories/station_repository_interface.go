package repositories

import "air-pollution/modules/core/models"

type StationRepositoryInterface interface {
	FindAll() ([]*models.Station, error)
}

type StationRepositorySupervisor struct {
	IsEnabled  bool
	Repository StationRepositoryInterface
}
