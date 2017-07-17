package file

import (
	"log"
	"main/tax_manager/utils"
	"fmt"
)

type csvHeaderStructure struct {
	municipality,
	date_from,
	date_to,
	tax_type,
	value int
}

func (this csvHeaderStructure) NewCSVHeaderStructure(header []string) (csvHeaderStructure) {
	for index, column := range header {
		switch column {
		case "municipality":
			this.municipality = index
		case "date_from":
			this.date_from = index
		case "date_to":
			this.date_to = index
		case "tax_type":
			this.tax_type = index
		case "value":
			this.value = index
		default:
			utils.Error("Header column `%s` is not parsable", header[index])
		}
	}
	log.Println(fmt.Sprintf("Parsed CVS file header indexes `%s`: ", this))
	return this
}
