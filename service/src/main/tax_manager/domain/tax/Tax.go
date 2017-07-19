package tax

import (
	"time"
	"fmt"
	"errors"
	"main/tax_manager/utils"
)

type Tax struct {
	Id             int64
	MunicipalityId int64
	From           time.Time
	To             time.Time
	TaxType        TaxType
	Value          float64
}

func NewTax(id int64, municipalityId int64, from time.Time, to time.Time, taxType TaxType, value float64) (Tax) {
	if municipalityId == 0 {
		utils.Error("Municipality id has to be set for tax")
	}

	if from.IsZero() {
		utils.Error("Tax from value has to be set for tax")
	}

	if to.IsZero() {
		utils.Error("Tax to value has to be set for tax")
	}

	if len(taxType) == 0 {
		utils.Error("Tax Type has to be set for tax")
	}

	if value == 0 {
		utils.Error("Tax value has to be set for tax")
	}

	return Tax{Id: id, MunicipalityId: municipalityId, From: from, To: to, TaxType: taxType, Value: value}
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
