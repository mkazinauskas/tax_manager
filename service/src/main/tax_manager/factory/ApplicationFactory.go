package factory

import (
	"main/tax_manager/domain/tax"
	"main/tax_manager/domain/municipality"
)

type ApplicationFactory interface {
	TaxRepository() (tax.TaxRepository)
	MunicipalityRepository() (municipality.MunicipalityRepository)
}
