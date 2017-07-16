package commands

import (
	"testing"
	"time"
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
)

func TestCalculateTax(testContext *testing.T) {
	currentTime, _ := time.Parse("yyyy-MM-dd", "2016-01-01")

	municipalityRepository := municipality.NewStubMunicipalityRepository(municipality.NewMunicipality(1, "Vilnius"))
	taxRepository := tax.NewStubTaxRepository()

	taxRate := NewCalculateTax(municipalityRepository, taxRepository).Calculate("Vilnius", currentTime)
	if taxRate != 0.1 {
		testContext.Error("Vilnius tax for 2016-01-01 was not 0.1")
	}
}
