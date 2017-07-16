package commands

import (
	"testing"
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
	"main/tax_manager/utils"
)


func TestCalculateNoTax(testContext *testing.T) {
	taxCalculationTime := utils.Parse("2015.01.01")

	municipalityRepository := municipality.NewStubMunicipalityRepository(municipality.NewMunicipality(1, "Vilnius"))

	taxRepository := tax.NewStubTaxRepository(tax.Tax{
		From:    utils.Parse("2016.01.01"),
		To:      utils.Parse("2016.12.31"),
		TaxType: tax.YEARLY,
		Value:   0.1,
	})

	taxRate := NewCalculateTax(municipalityRepository, taxRepository).Calculate("Vilnius", taxCalculationTime)
	if taxRate != 0.0 {
		testContext.Errorf("Vilnius tax for `%s` was `%f`, but expected 0.0", taxCalculationTime, taxRate)
	}
}

func TestCalculateYearlyTax(testContext *testing.T) {
	taxCalculationTime := utils.Parse("2016.01.01")

	municipalityRepository := municipality.NewStubMunicipalityRepository(municipality.NewMunicipality(1, "Vilnius"))

	taxRepository := tax.NewStubTaxRepository(tax.Tax{
		From:    utils.Parse("2016.01.01"),
		To:      utils.Parse("2016.12.31"),
		TaxType: tax.YEARLY,
		Value:   0.1,
	})

	taxRate := NewCalculateTax(municipalityRepository, taxRepository).Calculate("Vilnius", taxCalculationTime)
	if taxRate != 0.1 {
		testContext.Errorf("Vilnius tax for `%s` was `%f`, but expected 0.1", taxCalculationTime, taxRate)
	}
}
