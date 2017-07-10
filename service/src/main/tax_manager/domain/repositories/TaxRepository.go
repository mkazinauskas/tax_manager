package repositories

import (
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
	_ "main/tax_manager/domain"
	"main/tax_manager/domain"
	//"main/tax_manager/datasource"
)

func Save(tax domain.Tax) {
	//datasource.DB.Prepare()
	fmt.Printf("Tax id = %s", tax.Id)

}

func FindTaxByMunicipalityName(name string){
	//fmt.Printf()
}
