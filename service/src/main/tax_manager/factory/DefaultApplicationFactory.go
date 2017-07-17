package factory

import (
	"main/tax_manager/domain/tax"
	"main/tax_manager/domain/municipality"
)

type DefaultApplicationFactory struct {
}

func (DefaultApplicationFactory) TaxRepository() (tax.TaxRepository) {
	return tax.NewMySQLTaxRepository()
}

func (DefaultApplicationFactory) MunicipalityRepository() (municipality.MunicipalityRepository) {
	return municipality.NewMySQLMunicipalityRepository()
}

