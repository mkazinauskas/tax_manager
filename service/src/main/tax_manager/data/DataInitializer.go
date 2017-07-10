package data

import (
	"main/tax_manager/domain"
	"main/tax_manager/domain/repositories"
)

func Do(){
	repositories.Save(domain.Tax{Id:15, Value: 0.4})
}