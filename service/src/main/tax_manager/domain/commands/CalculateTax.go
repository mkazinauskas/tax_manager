package commands

import (
	"time"
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
	"github.com/aodin/date"
	"main/tax_manager/utils"
)

var municipalityRepository = municipality.MunicipalityRepository{}
var taxRepository = tax.TaxRepository{}

type CalculateTax struct {
}

func (this CalculateTax) Calculate(municipalityName string, date time.Time) (float64) {
	foundMunicipality := municipalityRepository.FindByName(municipalityName)
	if foundMunicipality == nil {
		utils.Error("No municipality was found by name = `%s`", municipalityName)
	}
	var calculatedTax float64 = 0
	for _, taxType := range []tax.TaxType{tax.DAILY, tax.WEEKLY, tax.MONTHLY, tax.YEARLY} {
		taxes := taxRepository.FindTaxByMunicipalityIdAndTaxType(foundMunicipality.Id, taxType)
		calculatedTax = this.getTax(taxes, date)
		if calculatedTax != 0 {
			break
		}
	}
	return calculatedTax
}

func (CalculateTax) getTax(taxes []tax.Tax, taxDate time.Time) (float64) {
	taxDateRange := date.NewRange(date.FromTime(taxDate), date.FromTime(taxDate))
	for _, tax := range taxes {
		dateRange := date.NewRange(date.FromTime(tax.From), date.FromTime(tax.To))
		if dateRange.Contains(taxDateRange) {
			return tax.Value
		}
	}
	return 0
}
