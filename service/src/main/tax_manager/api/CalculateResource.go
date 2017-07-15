package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"encoding/json"
	"time"
	"main/tax_manager"
)

func CalculateTax(w http.ResponseWriter, r *http.Request, rp httprouter.Params) {
	queryValues := r.URL.Query()
	municipalityName := queryValues.Get("municipalityName")
	if &municipalityName == nil {
		fmt.Fprint(w, Marshal(ErrorResponse{Message: "No `municipalityName` as request param was not found"}))
		return
	}
	date := queryValues.Get("date")
	if &date == nil {
		fmt.Fprint(w, Marshal(ErrorResponse{Message: "No `date` as request param was not found"}))
		return
	} else {
		_, parsingError := time.Parse(tax_manager.DEFAULT_DATE_FORMAT, date)
		if parsingError != nil {
			marshaled, _ := json.Marshal(ErrorResponse{Message: "Parameter `date` is not in format" + tax_manager.DEFAULT_DATE_FORMAT})
			fmt.Fprint(w, marshaled)
			return
		}

		fmt.Fprint(w, Marshal(CalculationResult{Tax: 0.1}))
	}
}
