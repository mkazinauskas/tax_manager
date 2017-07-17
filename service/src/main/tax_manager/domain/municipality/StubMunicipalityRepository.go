package municipality

type stubMunicipalityRepository struct {
	municipality Municipality
}

func NewStubMunicipalityRepository(municipality Municipality) (MunicipalityRepository) {
	return stubMunicipalityRepository{municipality: municipality}
}

func (this stubMunicipalityRepository) Save(municipality Municipality) (*Municipality) {
	return &this.municipality
}

func (this stubMunicipalityRepository) FindByName(name string) (*Municipality) {
	return &this.municipality
}

func (this stubMunicipalityRepository) FindById(id int64) (*Municipality) {
	return &this.municipality
}

func (this stubMunicipalityRepository) FindAll() ([]Municipality) {
	return []Municipality{this.municipality}
}

func (this stubMunicipalityRepository) DeleteAll() {
}

func (this stubMunicipalityRepository) Delete(municipality Municipality) {
}
