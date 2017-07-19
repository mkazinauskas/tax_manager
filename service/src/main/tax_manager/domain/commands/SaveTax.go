package commands

import (
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
	"main/tax_manager/factory"
	"log"
	"main/tax_manager/utils"
)

type saveTax struct {
	taxToSave              tax.Tax
	municipalityRepository municipality.MunicipalityRepository
	taxRepository          tax.TaxRepository
}

func NewSaveTax(taxToSave tax.Tax, factory factory.ApplicationFactory) (saveTax) {
	if taxToSave.MunicipalityId == 0 {
		utils.Error("Municipality id has to be set for tax")
	}

	if taxToSave.From.IsZero() {
		utils.Error("Tax from value has to be set for tax")
	}

	if taxToSave.To.IsZero() {
		utils.Error("Tax to value has to be set for tax")
	}

	if taxToSave.Value == 0 {
		utils.Error("Tax value has to be set for tax")
	}

	if len(taxToSave.TaxType) == 0 {
		utils.Error("Tax Type has to be set for tax")
	}

	return saveTax{
		taxToSave:              taxToSave,
		municipalityRepository: factory.MunicipalityRepository(),
		taxRepository:          factory.TaxRepository(),
	}
}

func (this saveTax) Handle() {
	savedMunicipality := this.municipalityRepository.FindById(this.taxToSave.MunicipalityId)
	if savedMunicipality == nil {
		utils.Error("Municipality not found by id `%d`", this.taxToSave.MunicipalityId)
		return
	}

	if this.taxRepository.IsExistingTax(this.taxToSave) {
		log.Println("Such tax already exsit", this.taxToSave)
		return
	}
	this.taxRepository.Save(this.taxToSave)
}
