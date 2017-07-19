package commands

import (
	"testing"
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
	"main/tax_manager/utils"
)

func TestCalculateNoTax(testContext *testing.T) {
	taxCalculationTime := utils.ParseDate("2015.01.01")

	municipalityRepository := municipality.NewStubMunicipalityRepository(*municipality.NewMunicipality(1, "Vilnius"))

	taxRepository := tax.NewStubTaxRepository(
		tax.NewTax(0,
			1,
			utils.ParseDate("2016.01.01"),
			utils.ParseDate("2016.12.31"),
			tax.YEARLY,
			0.1))

	taxRate := NewCalculateTax(municipalityRepository, taxRepository).Calculate("Vilnius", taxCalculationTime)
	if taxRate != 0.0 {
		testContext.Errorf("Vilnius tax for `%s` was `%f`, but expected 0.0", taxCalculationTime, taxRate)
	}
}

func TestCalculateYearlyTax(testContext *testing.T) {
	taxCalculationTime := utils.ParseDate("2016.01.01")

	municipalityRepository := municipality.NewStubMunicipalityRepository(*municipality.NewMunicipality(1, "Vilnius"))

	taxRepository := tax.NewStubTaxRepository(
		tax.NewTax(0,
			1,
			utils.ParseDate("2016.01.01"),
			utils.ParseDate("2016.12.31"),
			tax.YEARLY,
			0.2))

	taxRate := NewCalculateTax(municipalityRepository, taxRepository).Calculate("Vilnius", taxCalculationTime)
	if taxRate != 0.2 {
		testContext.Errorf("Vilnius tax for `%s` was `%f`, but expected 0.1", taxCalculationTime, taxRate)
	}
}

func TestCalculateMonthlyTax(testContext *testing.T) {
	taxCalculationTime := utils.ParseDate("2016.05.02")

	municipalityRepository := municipality.NewStubMunicipalityRepository(*municipality.NewMunicipality(1, "Vilnius"))

	taxRepository := tax.NewStubTaxRepository(
		tax.NewTax(0,
			1,
			utils.ParseDate("2016.05.01"),
			utils.ParseDate("2016.05.31"),
			tax.MONTHLY,
			0.4))

	taxRate := NewCalculateTax(municipalityRepository, taxRepository).Calculate("Vilnius", taxCalculationTime)
	if taxRate != 0.4 {
		testContext.Errorf("Vilnius tax for `%s` was `%f`, but expected 0.4", taxCalculationTime, taxRate)
	}
}

func TestCalculateDailyTax(testContext *testing.T) {
	taxCalculationTime := utils.ParseDate("2016.01.01")

	municipalityRepository := municipality.NewStubMunicipalityRepository(*municipality.NewMunicipality(1, "Vilnius"))

	taxRepository := tax.NewStubTaxRepository(
		tax.NewTax(0,
			1,
			utils.ParseDate("2016.01.01"),
			utils.ParseDate("2016.01.01"),
			tax.DAILY,
			0.1))

	taxRate := NewCalculateTax(municipalityRepository, taxRepository).Calculate("Vilnius", taxCalculationTime)
	if taxRate != 0.1 {
		testContext.Errorf("Vilnius tax for `%s` was `%f`, but expected 0.1", taxCalculationTime, taxRate)
	}
}
