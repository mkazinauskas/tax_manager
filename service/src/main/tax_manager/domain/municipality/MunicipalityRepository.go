package municipality

type MunicipalityRepository interface {
	Save(municipality Municipality) (*Municipality)
	FindByName(name string) (*Municipality)
	FindById(id int64) (*Municipality)
	FindAll() ([]Municipality)
	DeleteAll()
	Delete(municipality Municipality)
}
