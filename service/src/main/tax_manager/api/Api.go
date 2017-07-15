package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

func Initialize() {
	router := httprouter.New()
	router.GET("/", GetIndex)
	router.GET("/municipalities", GetAllMunicipalities)
	router.GET("/municipalities/:id", GetMunicipalityById)
	router.GET("/municipalities/:id/taxes", GetAllTaxes)
	router.GET("/calculate-tax", CalculateTax)

	log.Fatal(http.ListenAndServe(":8080", router))
}
