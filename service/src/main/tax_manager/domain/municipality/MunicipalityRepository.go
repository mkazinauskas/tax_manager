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

func (this MunicipalityRepository) Save(municipality Municipality) (Municipality) {
	result := this.database.Execute("INSERT `MUNICIPALITIES` SET name=?", municipality.Name)
	id, err := result.LastInsertId()
	utils.Check(err)

	return Municipality{Id: id, Name: municipality.Name}
}

func (this MunicipalityRepository) FindByName(name string) (*Municipality) {
	result := this.database.Query("SELECT * FROM `MUNICIPALITIES` WHERE name=?", name)
	return mapTo(result)
}

func mapTo(result *sql.Rows) (*Municipality) {
	if result.Next() {
		var id int64
		var name string
		result.Scan(&id, &name)
		return &Municipality{Id: id, Name: name}
	} else {
		return nil
	}
}
