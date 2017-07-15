package commands

import (
	"time"
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
)

var municipalityRepository = municipality.MunicipalityRepository{}
var taxRepository = tax.TaxRepository{}

func Calculate(municipality string, date time.Time) (rate float32) {
	//foundMunicipality := municipalityRepository.FindByName(municipality)
//	foundTaxes := taxRepository.FindTaxByMunicipalityIdAndTaxType(foundMunicipality.Id)

	return
}
