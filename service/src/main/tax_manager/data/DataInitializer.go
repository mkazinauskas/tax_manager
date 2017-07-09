package data

import "main/tax_manager/domain"

func Do(){
	taxRepo.Save(domain.Tax{15, "Test"})
}