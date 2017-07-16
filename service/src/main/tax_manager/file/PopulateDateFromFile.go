package file

import (
	"encoding/csv"
	"strings"
	"fmt"
	"io/ioutil"
	"errors"
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

type CSVHeaderStructure struct {
	municipality,
	date_from,
	date_to,
	tax_type,
	value int
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
	header := CSVHeaderStructure{}.populate(rows[0])

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
			this.applicationFactory.TaxRepository()).Save()
	}
}

func (populateDataFromFile) parseTax(row []string, header CSVHeaderStructure) (tax.Tax) {
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

func (this CSVHeaderStructure) populate(header []string) (CSVHeaderStructure) {
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
