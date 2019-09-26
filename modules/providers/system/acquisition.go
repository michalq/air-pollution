package system

import (
	"air-pollution/modules/core/models"
	"github.com/go-pg/pg/v9"
)

type AcquisitionRepository struct {
	db *pg.DB
}

func NewAcquisitionRepository(db *pg.DB) *AcquisitionRepository {
	return &AcquisitionRepository{db: db}
}

func (a *AcquisitionRepository) Save(acquisition *models.Acquisition) error {
	return nil
}
