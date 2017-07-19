package commands

import (
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
	"main/tax_manager/factory"
	"main/tax_manager/utils"
)

type saveMunicipality struct {
	municipalityToSave     municipality.Municipality
	municipalityRepository municipality.MunicipalityRepository
	taxRepository          tax.TaxRepository
}

func NewSaveMunicipality(municipalityToSave municipality.Municipality,
	factory factory.ApplicationFactory) (saveMunicipality) {
	if len(municipalityToSave.Name) == 0 {
		utils.Error("Municipality name is not present")
	}

	return saveMunicipality{
		municipalityToSave:     municipalityToSave,
		municipalityRepository: factory.MunicipalityRepository(),
		taxRepository:          factory.TaxRepository(),
	}
}

func (this saveMunicipality) Handle() (*municipality.Municipality) {

	savedMunicipality := this.municipalityRepository.FindByName(this.municipalityToSave.Name)
	if savedMunicipality == nil {
		savedMunicipality = this.municipalityRepository.Save(this.municipalityToSave)
	}
	return savedMunicipality
}
