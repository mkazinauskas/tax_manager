package tax

import (
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
)

func Save(tax Tax) {
	//datasource.Execute("INSERT municipalities SET municipality=? ",tax.Municipality, )

	fmt.Printf("Tax id = %s", tax.Id)

}

func FindTaxByMunicipalityName(name string){
	//fmt.Printf()
}
