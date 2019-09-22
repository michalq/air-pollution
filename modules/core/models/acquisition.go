package models

import (
	"time"
)

type Acquisition struct {
	Type     Type
	DateFrom time.Time
	DateTo   time.Time
	Value    string
}
