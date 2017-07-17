package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"main/tax_manager/domain/tax"
	"strconv"
	"main/tax_manager/utils"
	"encoding/json"
	"time"
	"main/tax_manager"
	"main/tax_manager/factory"
)

func GetAllTaxes(factory factory.ApplicationFactory) (httprouter.Handle) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		municipalityId, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: err.Error()}))
			return
		}
		utils.Check(err)

		foundMunicipality := factory.MunicipalityRepository().FindById(municipalityId)
		if foundMunicipality == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		taxes := factory.TaxRepository().FindTaxByMunicipalityId(municipalityId)
		fmt.Fprint(w, Marshal(taxes))
	}
}

func GetTaxById(factory factory.ApplicationFactory) (httprouter.Handle) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		municipalityId, municipalityIdError := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if municipalityIdError != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: municipalityIdError.Error()}))
			return
		}

		taxId, taxIdError := strconv.ParseInt(ps.ByName("tax-id"), 10, 64)
		if taxIdError != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: taxIdError.Error()}))
			return
		}

		foundMunicipality := factory.MunicipalityRepository().FindById(municipalityId)
		if foundMunicipality == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		foundTax := factory.TaxRepository().FindTaxByMunicipalityIdAndTaxId(municipalityId, taxId)
		if foundTax == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Fprint(w, Marshal(foundTax))
	}
}

func DeleteTaxById(factory factory.ApplicationFactory) (httprouter.Handle) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		municipalityId, municipalityIdError := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if municipalityIdError != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: municipalityIdError.Error()}))
			return
		}

		taxId, taxIdError := strconv.ParseInt(ps.ByName("tax-id"), 10, 64)
		if taxIdError != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: taxIdError.Error()}))
			return
		}

		foundMunicipality := factory.MunicipalityRepository().FindById(municipalityId)
		if foundMunicipality == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		foundTax := factory.TaxRepository().FindTaxByMunicipalityIdAndTaxId(municipalityId, taxId)
		if foundTax == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		factory.TaxRepository().Delete(*foundTax)

		w.WriteHeader(http.StatusNoContent)
	}
}

type SaveTaxRequest struct {
	From    string
	To      string
	TaxType string
	Value   float64
}

func SaveNewMunicipalityTax(factory factory.ApplicationFactory) (httprouter.Handle) {
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

		saveTaxRequest := SaveTaxRequest{}
		unmarshalError := json.NewDecoder(r.Body).Decode(&saveTaxRequest)
		utils.Check(unmarshalError)

		existingTax := factory.TaxRepository().FindTaxByMunicipalityIdAndTaxType(municipalityId, tax.FindTaxTypeByValue(saveTaxRequest.TaxType))
		if existingTax != nil {
			w.WriteHeader(http.StatusConflict)
			return
		}

		fromTime, fromTimeParseError := time.Parse(tax_manager.DEFAULT_DATE_FORMAT, saveTaxRequest.From)
		if fromTimeParseError != nil {
			fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: fmt.Sprintf("Property `from` has to be in format `%s`", tax_manager.DEFAULT_DATE_FORMAT)}))
		}

		toTime, toTimeParseError := time.Parse(tax_manager.DEFAULT_DATE_FORMAT, saveTaxRequest.To)
		if toTimeParseError != nil {
			fmt.Fprint(w, Marshal(ErrorResponse{ErrorMessage: fmt.Sprintf("Property `to` has to be in format `%s`", tax_manager.DEFAULT_DATE_FORMAT)}))
		}

		factory.TaxRepository().Save(
			tax.Tax{
				MunicipalityId: municipalityId,
				From:           fromTime,
				To:             toTime,
				TaxType:        tax.FindTaxTypeByValue(saveTaxRequest.TaxType),
				Value:          saveTaxRequest.Value})

		w.WriteHeader(http.StatusCreated)
	}
}
