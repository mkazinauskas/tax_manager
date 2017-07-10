package domain

import (
	"time"
)

type Tax struct {
	Id       int
	Municipality Municipality
	Duration time.Duration
	Value    float64
}
