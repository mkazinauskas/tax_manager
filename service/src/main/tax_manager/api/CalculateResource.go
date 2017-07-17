package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"time"
	"main/tax_manager"
	"main/tax_manager/domain/commands"
	"log"
	"main/tax_manager/factory"
)

func CalculateTax(factory factory.ApplicationFactory) (httprouter.Handle) {
	municipalityRepository := factory.MunicipalityRepository()
	taxRepository := factory.TaxRepository()
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		queryValues := r.URL.Query()
		municipalityName := queryValues.Get("municipality-name")

		if &municipalityName == nil {
			fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: "No `municipality-name` as request param was not found"}))
			log.Println("No `municipality-name` as request param was not found")
			return
		}

		if municipalityRepository.FindByName(municipalityName) == nil {
			fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: fmt.Sprintf("No `municipality` found with name `%s`", municipalityName)}))
			log.Println("No `municipality-name` as request param was not found")
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
				municipalityRepository,
				taxRepository).Calculate(municipalityName, date)
			fmt.Fprint(w, Marshal(CalculationResult{Tax: result}))
			log.Println(CalculationResult{Tax: result})
		}
	}
}
