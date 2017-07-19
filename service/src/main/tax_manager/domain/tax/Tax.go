package tax

import (
	"time"
	"fmt"
	"errors"
	"main/tax_manager/utils"
	"github.com/aodin/date"
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

	validateDatesByTaxType(from, to, taxType)

	if value == 0 {
		utils.Error("Tax value has to be set for tax")
	}

	return Tax{Id: id, MunicipalityId: municipalityId, From: from, To: to, TaxType: taxType, Value: value}
}
func validateDatesByTaxType(from time.Time, to time.Time, taxType TaxType) {
	dateRange := date.NewRange(date.FromTime(from), date.FromTime(to))
	if (taxType == YEARLY && dateRange.Days() < 365 && dateRange.Days() > 366) {
		wrongPeriod(taxType)
	}
	if (taxType == MONTHLY && dateRange.Days() < 28 && dateRange.Days() > 31) {
		wrongPeriod(taxType)
	}
	if (taxType == WEEKLY && dateRange.Days() != 7) {
		wrongPeriod(taxType)
	}
	if (taxType == DAILY && dateRange.Days() != 1) {
		wrongPeriod(taxType)
	}
}

func wrongPeriod(taxType TaxType) {
	utils.Error("Wrong from and to date period for tax type %s", taxType)
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
