package system

import (
	"air-pollution/daq/core/models"
	"github.com/go-pg/pg/v9"
)

type AcquisitionRepository struct {
	db *pg.DB
}

func NewAcquisitionRepository(db *pg.DB) *AcquisitionRepository {
	return &AcquisitionRepository{db: db}
}

func (a *AcquisitionRepository) FindAllByStationID(string) ([]*models.Acquisition, error) {
	return nil, nil
}

func (a *AcquisitionRepository) Persist(acquisition *models.Acquisition) error {
	return nil
}
