package system

import "github.com/go-pg/pg/v9"

type StationRepository struct {
	db *pg.DB
}

func NewStationRepository(db *pg.DB) *StationRepository {
	return &StationRepository{db: db}
}
