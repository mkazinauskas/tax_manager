package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"main/tax_manager/domain/municipality"
)

func GetAllMunicipalities(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	municipalities := municipality.MunicipalityRepository{}.FindAll()
	fmt.Fprint(w, Marshal(municipalities))
}
