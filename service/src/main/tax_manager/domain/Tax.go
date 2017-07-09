package domain

import (
	"time"
)

type Tax struct {
	Id       int
	Duration time.Duration
	Value    int
}
