package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"main/tax_manager/domain/tax"
	"strconv"
	"main/tax_manager/utils"
)

func GetAllTaxes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	value, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	utils.Check(err)

	taxes := tax.TaxRepository{}.FindTaxByMunicipalityId(value)
	fmt.Fprint(w, Marshal(taxes))
}
