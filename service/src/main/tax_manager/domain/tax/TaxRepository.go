package tax

import (
	_ "github.com/ziutek/mymysql/godrv"
	"main/tax_manager/datasource"
	"database/sql"
	"time"
)

type TaxRepository struct {
	database datasource.Database
}

func (this TaxRepository) Save(tax Tax) {
	this.database.Execute("INSERT `TAXES` SET `MUNICIPALITY_ID`=?,`FROM`=?,`TO`=?,`TAX_TYPE`=?,`VALUE`=?",
		tax.MunicipalityId,
		tax.From.Format("yyyy-MM-dd"),
		tax.To.Format("yyyy-MM-dd"),
		tax.TaxType,
		tax.Value)
}

func (this *TaxRepository) FindTaxByMunicipalityId(id int64) (*Tax) {
	result := this.database.Query("SELECT * FROM `TAXES` WHERE MUNICIPALITY_ID=?", id)
	return mapTo(result)
}

func mapTo(result *sql.Rows) (*Tax) {
	if result.Next() {
		var id int64
		var municipalityId int64
		var from string
		var to string
		var taxType string
		var value float64
		result.Scan(&id, &municipalityId, &from, &to, &taxType, &value)

		parsedFrom, _ := time.Parse("yyyy-MM-dd", from)
		parsedTo, _ := time.Parse("yyyy-MM-dd", to)
		return &Tax{
			Id:             id,
			MunicipalityId: municipalityId,
			From:           parsedFrom,
			To:             parsedTo,
			TaxType:        FindTaxTypeByValue(taxType),
			Value:          value}
	} else {
		return nil
	}
}
