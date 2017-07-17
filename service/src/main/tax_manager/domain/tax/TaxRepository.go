package tax

type TaxRepository interface {
	Save(tax Tax)
	IsExistingTax(tax Tax)(bool)
	FindTaxByMunicipalityIdAndTaxType(id int64, taxType TaxType) ([]Tax)
	FindTaxByMunicipalityIdAndTaxId(municipalityId int64, taxId int64) (*Tax)
	FindTaxByMunicipalityId(id int64) ([]Tax)
	DeleteAll()
	Delete(tax Tax)
}