package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"main/tax_manager/domain/tax"
	"strconv"
	"main/tax_manager/utils"
	"main/tax_manager/domain/municipality"
)

func GetAllTaxes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	value, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	utils.Check(err)

	taxes := tax.TaxRepository{}.FindTaxByMunicipalityId(value)
	fmt.Fprint(w, Marshal(taxes))
}

func GetTaxById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	municipalityId, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	utils.Check(err)

	taxId, err := strconv.ParseInt(ps.ByName("tax-id"), 10, 64)
	utils.Check(err)

	foundMunicipality := municipality.MunicipalityRepository{}.FindById(municipalityId)
	if foundMunicipality == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	foundTax := tax.TaxRepository{}.FindTaxByMunicipalityIdAndTaxId(municipalityId, taxId)
	if foundTax == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, Marshal(foundTax))
}
