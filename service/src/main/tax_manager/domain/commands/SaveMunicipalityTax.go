package commands

import (
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
)

type saveMunicipalityAndTax struct {
	municipalityToSave     municipality.Municipality
	taxToSave              tax.Tax
	municipalityRepository municipality.MunicipalityRepository
	taxRepository          tax.TaxRepository
}

func NewSaveMunicipalityAndTax(municipalityToSave municipality.Municipality,
	taxToSave tax.Tax,
	municipalityRepository municipality.MunicipalityRepository,
	taxRepository tax.TaxRepository) (saveMunicipalityAndTax) {
	return saveMunicipalityAndTax{
		municipalityToSave:     municipalityToSave,
		taxToSave:              taxToSave,
		municipalityRepository: municipalityRepository,
		taxRepository:          taxRepository,
	}
}

func (this saveMunicipalityAndTax) Save() {
	savedMunicipality := this.municipalityRepository.FindByName(this.municipalityToSave.Name)
	if savedMunicipality == nil {
		savedMunicipality = this.municipalityRepository.Save(this.municipalityToSave)
	}

	this.taxToSave.MunicipalityId = savedMunicipality.Id
	existingTaxes := this.taxRepository.FindTaxByMunicipalityIdAndTaxType(savedMunicipality.Id, this.taxToSave.TaxType)
	if existingTaxes == nil {
		this.taxRepository.Save(this.taxToSave)
	}
}
