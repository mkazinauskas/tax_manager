package municipality

import (
	_ "github.com/ziutek/mymysql/godrv"
	"main/tax_manager/datasource"
	"main/tax_manager/utils"
	"database/sql"
)

func Save(municipality Municipality) (Municipality) {
	result := datasource.Execute("INSERT municipalities SET name=?", municipality.Name)
	id, err := result.LastInsertId()
	utils.Check(err)

	return Municipality{Id: id, Name: municipality.Name}
}

func FindByName(name string) (*Municipality) {
	result := datasource.Query("SELECT * FROM municipalities WHERE name=?", name)
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
