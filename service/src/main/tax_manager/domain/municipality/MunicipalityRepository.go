package municipality

import (
	_ "github.com/ziutek/mymysql/godrv"
	"main/tax_manager/datasource"
	"main/tax_manager/utils"
	"database/sql"
)

type MunicipalityRepository struct {
	database datasource.Database
}

func (this MunicipalityRepository) Save(municipality Municipality) (*Municipality) {
	result := this.database.Execute("INSERT `MUNICIPALITIES` SET name=?", municipality.Name)
	id, err := result.LastInsertId()
	utils.Check(err)

	return &Municipality{Id: id, Name: municipality.Name}
}

func (this MunicipalityRepository) FindByName(name string) (*Municipality) {
	result := this.database.Query("SELECT * FROM `MUNICIPALITIES` WHERE name=?", name)
	return this.takeFirst(mapTo(result))
}

func (this MunicipalityRepository) FindById(id int64) (*Municipality) {
	result := this.database.Query("SELECT * FROM `MUNICIPALITIES` WHERE id=?", id)
	return this.takeFirst(mapTo(result))
}

func (MunicipalityRepository) takeFirst(municipalities []Municipality) (*Municipality) {
	if len(municipalities) == 1 {
		return &municipalities[0]
	} else {
		return nil
	}
}

func (this MunicipalityRepository) FindAll() ([]Municipality) {
	result := this.database.Query("SELECT * FROM `MUNICIPALITIES`")
	return mapTo(result)
}

func (this MunicipalityRepository) DeleteAll() {
	this.database.Query("DELETE FROM `MUNICIPALITIES`")
}

func (this MunicipalityRepository) Delete(municipality Municipality) {
	this.database.Query("DELETE FROM `MUNICIPALITIES` WHERE `ID`=?", municipality.Id)
}

func mapTo(result *sql.Rows) ([]Municipality) {
	municipalities := []Municipality{}
	for result.Next() {
		var id int64
		var name string
		result.Scan(&id, &name)
		municipalities = append(municipalities, Municipality{Id: id, Name: name})
	}
	return municipalities
}
