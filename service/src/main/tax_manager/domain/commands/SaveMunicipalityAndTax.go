package commands

import (
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
	"main/tax_manager/factory"
	"log"
)

type saveMunicipalityAndTax struct {
	municipalityToSave     municipality.Municipality
	taxToSave              tax.Tax
	municipalityRepository municipality.MunicipalityRepository
	taxRepository          tax.TaxRepository
}

func NewSaveMunicipalityAndTax(municipalityToSave municipality.Municipality,
	taxToSave tax.Tax,
	factory factory.ApplicationFactory) (saveMunicipalityAndTax) {

	return saveMunicipalityAndTax{
		municipalityToSave:     municipalityToSave,
		taxToSave:              taxToSave,
		municipalityRepository: factory.MunicipalityRepository(),
		taxRepository:          factory.TaxRepository(),
	}
}

func (this saveMunicipalityAndTax) Handle() {
	savedMunicipality := this.municipalityRepository.FindByName(this.municipalityToSave.Name)
	if savedMunicipality == nil {
		savedMunicipality = this.municipalityRepository.Save(this.municipalityToSave)
	}

	this.taxToSave.MunicipalityId = savedMunicipality.Id
	if this.taxRepository.IsExistingTax(this.taxToSave) {
		log.Println("Such tax already exsit", this.taxToSave)
		return
	}
	this.taxRepository.Save(this.taxToSave)
}
