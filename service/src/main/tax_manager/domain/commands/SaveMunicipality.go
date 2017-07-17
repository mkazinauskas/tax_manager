package commands

import (
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
	"main/tax_manager/factory"
)

type saveMunicipality struct {
	municipalityToSave     municipality.Municipality
	municipalityRepository municipality.MunicipalityRepository
	taxRepository          tax.TaxRepository
}

func NewSaveMunicipality(municipalityToSave municipality.Municipality,
	factory factory.ApplicationFactory) (saveMunicipality) {

	return saveMunicipality{
		municipalityToSave:     municipalityToSave,
		municipalityRepository: factory.MunicipalityRepository(),
		taxRepository:          factory.TaxRepository(),
	}
}

func (this saveMunicipality) Handle() (*municipality.Municipality){
	savedMunicipality := this.municipalityRepository.FindByName(this.municipalityToSave.Name)
	if savedMunicipality == nil {
		savedMunicipality = this.municipalityRepository.Save(this.municipalityToSave)
	}
	return savedMunicipality
}
