package file

import (
	"errors"
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
			panic(errors.New(fmt.Sprintf("Header column %s", header[index])))
		}
	}
	fmt.Println(this)
	return this
}
