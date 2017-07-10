package tax

import (
	"time"
	"main/tax_manager/domain/municipality"
)

type Tax struct {
	Id       int
	Municipality municipality.Municipality
	Duration time.Duration
	Value    float64
}
