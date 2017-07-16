package file

import (
	"encoding/csv"
	"strings"
	"fmt"
	"io/ioutil"
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
	"time"
	"main/tax_manager"
	"strconv"
	"main/tax_manager/utils"
	"main/tax_manager/domain/commands"
	"main/tax_manager/factory"
)

const DEFAULT_COLUMN_LENGTH = 5

type populateDataFromFile struct {
	applicationFactory factory.ApplicationFactory
}

func NewPopulateDataFromFile(applicationFactory factory.ApplicationFactory) (populateDataFromFile) {
	return populateDataFromFile{applicationFactory: applicationFactory}
}

func (this populateDataFromFile) Populate(filePath string) {
	contentAsBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(strings.NewReader(string(contentAsBytes)))

	rows, err := r.ReadAll()
	utils.Check(err)

	fmt.Println(rows)

	this.validateRowsLength(rows)
	fmt.Println(rows)
	header := csvHeaderStructure{}.NewCSVHeaderStructure(rows[0])

	fmt.Println(header)
	for _, row := range rows[1:] {
		fmt.Println(row)

		parsedMunicipality := municipality.NewMunicipality(0, row[header.municipality])
		fmt.Println(parsedMunicipality)

		parsedTax := this.parseTax(row, header)
		fmt.Println(parsedTax)

		commands.NewSaveMunicipalityAndTax(
			parsedMunicipality,
			parsedTax,
			this.applicationFactory.MunicipalityRepository(),
			this.applicationFactory.TaxRepository()).Handle()
	}
}

func (populateDataFromFile) parseTax(row []string, header csvHeaderStructure) (tax.Tax) {
	from, fromDateParsing := time.Parse(tax_manager.DEFAULT_DATE_FORMAT, row[header.date_from])
	utils.CheckError(fromDateParsing, "Failed to parse date from `%s`", string(row[header.date_from]))

	to, toDateParsing := time.Parse(tax_manager.DEFAULT_DATE_FORMAT, row[header.date_to])
	utils.CheckError(toDateParsing, "Failed to parse date to `%s`", row[header.date_to])

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
