package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"main/tax_manager/domain/municipality"
	"strconv"
	"main/tax_manager/utils"
	"encoding/json"
)

func GetAllMunicipalities(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	municipalities := municipality.MunicipalityRepository{}.FindAll()
	fmt.Fprint(w, Marshal(municipalities))
}

type SaveNewMunicipalityRequest struct {
	Name string
}

func SaveNewMunicipality(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	saveMunicipalityRequest := SaveNewMunicipalityRequest{}
	unmarshalError := json.NewDecoder(r.Body).Decode(&saveMunicipalityRequest)
	utils.Check(unmarshalError)

	existingMunicipality := municipality.MunicipalityRepository{}.FindByName(saveMunicipalityRequest.Name)
	if existingMunicipality != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	municipality.MunicipalityRepository{}.Save(municipality.Municipality{Name: saveMunicipalityRequest.Name})

	w.WriteHeader(http.StatusCreated)
}

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

func DeleteMunicipalityById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	value, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	utils.Check(err)

	foundMunicipality := municipality.MunicipalityRepository{}.FindById(value)
	if foundMunicipality == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	municipality.MunicipalityRepository{}.Delete(*foundMunicipality)
	w.WriteHeader(http.StatusNoContent)
}
