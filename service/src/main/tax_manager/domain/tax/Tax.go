package tax

import (
	"time"
	"fmt"
	"errors"
)

type Tax struct {
	Id             int64
	MunicipalityId int64
	From           time.Time
	To             time.Time
	TaxType        TaxType
	Value          float64
}

type TaxType string

const (
	YEARLY  TaxType = "yearly"
	MONTHLY TaxType = "monthly"
	WEEKLY  TaxType = "weekly"
	DAILY   TaxType = "daily"
)

func FindTaxTypeByValue(value string) (TaxType) {
	switch value {
	case string(YEARLY):
		return YEARLY
	case string(MONTHLY):
		return MONTHLY
	case string(WEEKLY):
		return WEEKLY
	case string(DAILY):
		return DAILY
	default:
		panic(errors.New(fmt.Sprintf("Tax type not found by value %s", value)))
	}
}
