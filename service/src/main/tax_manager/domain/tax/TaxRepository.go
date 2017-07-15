package tax

import (
	_ "github.com/ziutek/mymysql/godrv"
	"main/tax_manager/datasource"
	"database/sql"
	"time"
	"main/tax_manager"
)

type TaxRepository struct {
	database datasource.Database
}

func (this TaxRepository) Save(tax Tax) {
	this.database.Execute("INSERT `TAXES` SET `MUNICIPALITY_ID`=?,`FROM`=?,`TO`=?,`TAX_TYPE`=?,`VALUE`=?",
		tax.MunicipalityId,
		tax.From.Format(tax_manager.DEFAULT_DATE_FORMAT),
		tax.To.Format(tax_manager.DEFAULT_DATE_FORMAT),
		string(tax.TaxType),
		tax.Value)
}

func (this *TaxRepository) FindTaxByMunicipalityId(id int64, taxType TaxType) (*[]Tax) {
	result := this.database.Query("SELECT * FROM `TAXES` WHERE `MUNICIPALITY_ID`=? AND `TAX_TYPE=?`", id, string(taxType))
	return mapTo(result)
}

func mapTo(result *sql.Rows) (*[]Tax) {
	foundTaxes := []Tax{}
	for result.Next() {
		var id int64
		var municipalityId int64
		var from string
		var to string
		var taxType string
		var value float64
		result.Scan(&id, &municipalityId, &from, &to, &taxType, &value)

		parsedFrom, _ := time.Parse(tax_manager.DEFAULT_DATE_FORMAT, from)
		parsedTo, _ := time.Parse(tax_manager.DEFAULT_DATE_FORMAT, to)
		foundTaxes = append(foundTaxes, Tax{
			Id:             id,
			MunicipalityId: municipalityId,
			From:           parsedFrom,
			To:             parsedTo,
			TaxType:        FindTaxTypeByValue(taxType),
			Value:          value})
	}
	return &foundTaxes

}
