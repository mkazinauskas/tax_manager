package file

import (
	"encoding/csv"
	"strings"
	"fmt"
	"io/ioutil"
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
	"strconv"
	"main/tax_manager/utils"
	"main/tax_manager/domain/commands"
	"main/tax_manager/factory"
	"log"
)

const DEFAULT_COLUMN_LENGTH = 5

type populateDataFromFile struct {
	applicationFactory factory.ApplicationFactory
}

func NewPopulateDataFromFile(applicationFactory factory.ApplicationFactory) (populateDataFromFile) {
	return populateDataFromFile{applicationFactory: applicationFactory}
}

func (this populateDataFromFile) Populate(filePath string) {
	log.Println(fmt.Sprintf("Reading file: `%s`", filePath))
	contentAsBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(strings.NewReader(string(contentAsBytes)))

	rows, err := r.ReadAll()
	utils.Check(err)

	log.Println(fmt.Sprintf("Rows from CSV file: `%s`", rows))

	this.validateRowsLength(rows)
	header := csvHeaderStructure{}.NewCSVHeaderStructure(rows[0])

	for _, row := range rows[1:] {
		log.Println(fmt.Sprintf("Parsing row `%s`", row))

		parsedMunicipality := municipality.NewMunicipality(0, row[header.municipality])
		log.Println(fmt.Sprintf("Parsed municipality `%s`", parsedMunicipality))

		parsedTax := this.parseTax(row, header)
		log.Println(fmt.Sprintf("Parsed tax `%s`", parsedTax))

		savedMunicipality :=commands.NewSaveMunicipality(parsedMunicipality, this.applicationFactory).Handle()

		parsedTax.MunicipalityId = savedMunicipality.Id

		commands.NewSaveTax(parsedTax, this.applicationFactory).Handle()
	}
}

func (populateDataFromFile) parseTax(row []string, header csvHeaderStructure) (tax.Tax) {
	from := utils.ParseDate(row[header.date_from])
	to := utils.ParseDate(row[header.date_to])

	value, floatValueError := strconv.ParseFloat(row[header.value], 64)
	utils.CheckError(floatValueError, "Failed to parse float value from `%s`", row[header.value])

	return tax.Tax{
		From:    from,
		To:      to,
		TaxType: tax.FindTaxTypeByValue(row[header.tax_type]),
		Value:   value}
}

func (this populateDataFromFile) validateRowsLength(rows [][]string) {
	for index, row := range rows {
		if len(row) != DEFAULT_COLUMN_LENGTH {
			utils.Error("Row `%s` column count is not %s", index, DEFAULT_COLUMN_LENGTH)
		}
	}
}
