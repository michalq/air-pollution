package system

import (
	"air-pollution/daq/models"
	"github.com/go-pg/pg/v9"
)

type StationRepository struct {
	db *pg.DB
}

func NewStationRepository(db *pg.DB) *StationRepository {
	return &StationRepository{db: db}
}

func (s *StationRepository) FindAll() ([]*models.Station, error) {
	return nil, nil
}
