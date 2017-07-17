package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"main/tax_manager/factory"
	"main/tax_manager"
)

func Initialize(factory factory.ApplicationFactory) {
	log.Println("Loading api mappings")
	router := httprouter.New()
	router.GET("/", GetIndex)
	router.GET("/municipalities", GetAllMunicipalities(factory))
	router.POST("/municipalities", SaveNewMunicipality(factory))
	router.GET("/municipalities/:id", GetMunicipalityById(factory))
	router.DELETE("/municipalities/:id", DeleteMunicipalityById(factory))
	router.GET("/municipalities/:id/taxes", GetAllTaxes(factory))
	router.POST("/municipalities/:id/taxes", SaveNewMunicipalityTax(factory))
	router.GET("/municipalities/:id/taxes/:tax-id", GetTaxById(factory))
	router.DELETE("/municipalities/:id/taxes/:tax-id", DeleteTaxById(factory))
	router.GET("/calculate-tax", CalculateTax(factory))

	log.Println("SERVER IS STARTED ON PORT", tax_manager.API_PORT)
	log.Fatal(http.ListenAndServe(tax_manager.API_PORT, router))
}
