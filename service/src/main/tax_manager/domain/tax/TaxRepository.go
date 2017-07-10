package tax

import (
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
)

func Save(tax Tax) {
	startTime, _ :=time.Parse(time.RFC3339, "2012-11-01T22:00:00+00:00")
	startTime.Year() = 5
	//datasource.Execute("INSERT municipalities SET municipality=? ",tax.Municipality, )
	fmt.Printf("Tax id = %s", tax.Id)

}

func FindTaxByMunicipalityName(name string){
	//fmt.Printf()
}
