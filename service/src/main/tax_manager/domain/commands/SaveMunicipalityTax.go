package commands

import (
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
)

type SaveMunicipalityAndTax struct {
	MunicipalityToSave municipality.Municipality
	TaxToSave          tax.Tax
}

func (this SaveMunicipalityAndTax) Save() {
	savedMunicipality := municipality.MunicipalityRepository{}.FindByName(this.MunicipalityToSave.Name)
	if savedMunicipality == nil {
		savedMunicipality = municipality.MunicipalityRepository{}.Save(this.MunicipalityToSave)
	}

	this.TaxToSave.MunicipalityId = savedMunicipality.Id
	taxRepository.Save(this.TaxToSave)
}