package tax

import (
	"time"
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

var taxTypes = map[string]TaxType{
	YEARLY:  YEARLY,
	MONTHLY: MONTHLY,
	WEEKLY:  WEEKLY,
	DAILY:   DAILY,
}

func FindByValue(value string) (TaxType) {
	return taxTypes[value]
}
