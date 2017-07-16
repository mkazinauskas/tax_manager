package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"time"
	"main/tax_manager"
	"main/tax_manager/domain/commands"
	"log"
	"main/tax_manager/domain/municipality"
	"main/tax_manager/domain/tax"
)

func CalculateTax(w http.ResponseWriter, r *http.Request, rp httprouter.Params) {
	queryValues := r.URL.Query()
	municipalityName := queryValues.Get("municipalityName")
	log.Println(municipalityName)
	if &municipalityName == nil {
		fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: "No `municipalityName` as request param was not found"}))
		log.Println("No `municipalityName` as request param was not found")
		return
	}

	if municipality.NewMySQLMunicipalityRepository().FindByName(municipalityName) == nil {
		fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: fmt.Sprintf("No `municipality` found with name `%s`", municipalityName)}))
		log.Println("No `municipalityName` as request param was not found")
		return
	}

	date := queryValues.Get("date")
	log.Println(date)
	if &date == nil {
		fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: "No `date` as request param was not found"}))
		log.Println("No `date` as request param was not found")
		return
	} else {
		date, parsingError := time.Parse(tax_manager.DEFAULT_DATE_FORMAT, date)
		if parsingError != nil {
			fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: "Parameter `date` is not in format " + tax_manager.DEFAULT_DATE_FORMAT}))
			return
		}

		result := commands.NewCalculateTax(
			municipality.NewMySQLMunicipalityRepository(),
			tax.NewMySQLTaxRepository()).Calculate(municipalityName, date)
		fmt.Fprint(w, Marshal(CalculationResult{Tax: result}))
		log.Println(CalculationResult{Tax: result})
	}
}
