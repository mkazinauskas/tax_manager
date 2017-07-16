package tax

import (
	"main/tax_manager/datasource"
)

type stubTaxRepository struct {
	database datasource.Database
}

func NewStubTaxRepository() (TaxRepository) {
	return mySQLTaxRepository{}
}

func (this stubTaxRepository) Save(tax Tax) {
}

func (this stubTaxRepository) FindTaxByMunicipalityIdAndTaxType(id int64, taxType TaxType) ([]Tax) {
	return []Tax{}
}

func (this stubTaxRepository) FindTaxByMunicipalityIdAndTaxId(municipalityId int64, taxId int64) (*Tax) {
	return &Tax{}
}

func (this stubTaxRepository) FindTaxByMunicipalityId(id int64) ([]Tax) {
	return []Tax{}
}

func (this stubTaxRepository) DeleteAll() {
}

func (this stubTaxRepository) Delete(tax Tax) {
}
