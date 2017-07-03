package domain

import (
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
)

type TaxRepository struct {
}

func (h *TaxRepository) Save(tax Tax) {
	fmt.Printf("Tax id = %s", tax.Id)

}
