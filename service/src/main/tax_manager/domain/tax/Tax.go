package tax

import (
	"time"
	"main/tax_manager/domain/municipality"
)

type Tax struct {
	Id           int
	Municipality municipality.Municipality
	Duration     time.Duration
	Type         TaxType
	Value        float64
}

type TaxType string

const (
	YEARLY  TaxType = "yearly"
	MONTHLY TaxType = "monthly"
	DAILY   TaxType = "daily"
)

var taxTypes = map[string]TaxType{
	YEARLY:  YEARLY,
	MONTHLY: MONTHLY,
	DAILY:   DAILY,
}

func (TaxType) findByValue(value string) (TaxType) {
	return taxTypes[value]
}
