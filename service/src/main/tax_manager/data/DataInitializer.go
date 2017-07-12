package data

import (
	"main/tax_manager/domain/municipality"
	"fmt"
)

func InitDefaultData() {
	createMunicipality("Vilnius")
	createMunicipality("Kaunas")
}

func createMunicipality(name string) {
	existingMunicipality := municipality.FindByName(name)
	fmt.Println(existingMunicipality)
	if existingMunicipality == nil {
		municipality.Save(municipality.Municipality{Name: name})
	}
}
