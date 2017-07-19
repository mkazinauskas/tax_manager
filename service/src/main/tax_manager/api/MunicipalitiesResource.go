package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"strconv"
	"main/tax_manager/utils"
	"encoding/json"
	"main/tax_manager/factory"
	"main/tax_manager/domain/commands"
	"main/tax_manager/domain/municipality"
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

		if len(saveMunicipalityRequest.Name) == 0 {
			fmt.Fprint(w, Marshal(
				ErrorResponse{ErrorMessage: fmt.Sprintf("Property `name` has to be set")}))
			return
		}

		existingMunicipality := factory.MunicipalityRepository().FindByName(saveMunicipalityRequest.Name)
		if existingMunicipality != nil {
			w.WriteHeader(http.StatusConflict)
			return
		}
		municipalityToSave := municipality.Municipality{Name: saveMunicipalityRequest.Name}
		commands.NewSaveMunicipality(municipalityToSave, factory).Handle()

		w.WriteHeader(http.StatusCreated)
		//Location header with id has to be added... Don't know how to add to headers....
	}
}

func GetMunicipalityById(factory factory.ApplicationFactory) (httprouter.Handle) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		municipalityId, municipalityIdError := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if municipalityIdError != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: municipalityIdError.Error()}))
			return
		}

		foundMunicipality := factory.MunicipalityRepository().FindById(municipalityId)
		if foundMunicipality == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Fprint(w, Marshal(foundMunicipality))
	}
}

func DeleteMunicipalityById(factory factory.ApplicationFactory) (httprouter.Handle) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		municipalityId, municipalityIdError := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if municipalityIdError != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: municipalityIdError.Error()}))
			return
		}

		foundMunicipality := factory.MunicipalityRepository().FindById(municipalityId)
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
