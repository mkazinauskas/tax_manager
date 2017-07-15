package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"main/tax_manager/domain/municipality"
	"fmt"
	"strconv"
	"main/tax_manager/utils"
)

func GetMunicipalityById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	value, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	utils.Check(err)

	foundMunicipality := municipality.MunicipalityRepository{}.FindById(value)
	if foundMunicipality == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, Marshal(foundMunicipality))
}
