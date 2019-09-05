package models

import (
	"github.com/michalq/go-gios-api-client/models"
	"time"
)

type Acquisition struct {
	Type     models.AcquisitionType
	DateFrom time.Time
	DateTo   time.Time
	Value    string
}
