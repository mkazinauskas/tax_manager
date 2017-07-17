package tax

type stubTaxRepository struct {
	stubbedTax Tax
}

func NewStubTaxRepository(stubbedTax Tax) (TaxRepository) {
	return stubTaxRepository{stubbedTax: stubbedTax}
}

func (this stubTaxRepository) Save(tax Tax) {
}

func (this stubTaxRepository) IsExistingTax(tax Tax)(bool){
	return false
}

func (this stubTaxRepository) FindTaxByMunicipalityIdAndTaxType(id int64, taxType TaxType) ([]Tax) {
	return []Tax{this.stubbedTax}
}

func (this stubTaxRepository) FindTaxByMunicipalityIdAndTaxId(municipalityId int64, taxId int64) (*Tax) {
	return &this.stubbedTax
}

func (this stubTaxRepository) FindTaxByMunicipalityId(id int64) ([]Tax) {
	return []Tax{this.stubbedTax}
}

func (this stubTaxRepository) DeleteAll() {
}

func (this stubTaxRepository) Delete(tax Tax) {
}
