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
	router.POST("/municipalities", SaveNewMunicipality)
	router.GET("/municipalities/:id", GetMunicipalityById)
	router.GET("/municipalities/:id/taxes", GetAllTaxes)
	router.GET("/municipalities/:id/taxes/:tax-id", GetTaxById)
	router.GET("/calculate-tax", CalculateTax)

	log.Fatal(http.ListenAndServe(":8080", router))
}
