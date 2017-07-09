package domain

import (
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
)

func Save(tax Tax) {
	fmt.Printf("Tax id = %s", tax.Id)

}
