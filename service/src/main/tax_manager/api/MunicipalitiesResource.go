package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"main/tax_manager/domain/municipality"
	"strconv"
	"main/tax_manager/utils"
	"encoding/json"
	"main/tax_manager/factory"
)

func GetAllMunicipalities(factory factory.ApplicationFactory) (httprouter.Handle) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		municipalities := factory.MunicipalityRepository().FindAll()
		fmt.Fprint(w, Marshal(municipalities))
	}
}

type SaveNewMunicipalityRequest struct {
	Name string
}

func SaveNewMunicipality(factory factory.ApplicationFactory) (httprouter.Handle) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		saveMunicipalityRequest := SaveNewMunicipalityRequest{}
		unmarshalError := json.NewDecoder(r.Body).Decode(&saveMunicipalityRequest)
		utils.Check(unmarshalError)

		existingMunicipality := factory.MunicipalityRepository().FindByName(saveMunicipalityRequest.Name)
		if existingMunicipality != nil {
			w.WriteHeader(http.StatusConflict)
			return
		}
		factory.MunicipalityRepository().Save(municipality.Municipality{Name: saveMunicipalityRequest.Name})

		w.WriteHeader(http.StatusCreated)
	}
}

func GetMunicipalityById(factory factory.ApplicationFactory) (httprouter.Handle) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		value, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
		utils.Check(err)

		foundMunicipality := factory.MunicipalityRepository().FindById(value)
		if foundMunicipality == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Fprint(w, Marshal(foundMunicipality))
	}
}

func DeleteMunicipalityById(factory factory.ApplicationFactory) (httprouter.Handle) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		value, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
		utils.Check(err)

		foundMunicipality := factory.MunicipalityRepository().FindById(value)
		if foundMunicipality == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		foundTaxes := factory.TaxRepository().FindTaxByMunicipalityId(foundMunicipality.Id)
		if len(foundTaxes) > 0 {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		factory.MunicipalityRepository().Delete(*foundMunicipality)
		w.WriteHeader(http.StatusNoContent)
	}
}
