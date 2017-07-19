package commands

import (
	"time"
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
	"github.com/aodin/date"
	"main/tax_manager/utils"
	"log"
)

type calculateTax struct {
	municipalityRepository municipality.MunicipalityRepository
	taxRepository          tax.TaxRepository
}

func NewCalculateTax(municipalityRepo municipality.MunicipalityRepository, taxRepo tax.TaxRepository) (calculateTax) {
	return calculateTax{municipalityRepository: municipalityRepo, taxRepository: taxRepo}
}

func (this calculateTax) Calculate(municipalityName string, date time.Time) (float64) {
	if len(municipalityName) == 0 {
		utils.Error("Municipality name is not present")
	}
	foundMunicipality := this.municipalityRepository.FindByName(municipalityName)
	if foundMunicipality == nil {
		utils.Error("No municipality was found by name = `%s`", municipalityName)
	}
	var calculatedTax float64 = 0
	for _, taxType := range []tax.TaxType{tax.DAILY, tax.WEEKLY, tax.MONTHLY, tax.YEARLY} {
		taxes := this.taxRepository.FindTaxByMunicipalityIdAndTaxType(foundMunicipality.Id, taxType)
		calculatedTax = this.getTax(taxes, date)
		if calculatedTax != 0 {
			break
		}
	}
	return calculatedTax
}

func (calculateTax) getTax(taxes []tax.Tax, taxDate time.Time) (float64) {
	taxDateRange := date.NewRange(date.FromTime(taxDate), date.FromTime(taxDate))
	for _, taxForEvaluation := range taxes {
		log.Println("from", taxForEvaluation.From)
		log.Println("to", taxForEvaluation.To)
		dateRange := date.NewRange(date.FromTime(taxForEvaluation.From), date.FromTime(taxForEvaluation.To))
		if dateRange.Contains(taxDateRange) {
			return taxForEvaluation.Value
		}
	}
	return 0
}
