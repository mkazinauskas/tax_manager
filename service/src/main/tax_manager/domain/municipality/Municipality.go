package municipality

import "main/tax_manager/utils"

type Municipality struct {
	Id   int64
	Name string
}

func NewMunicipality(id int64, name string) (Municipality) {
	if len(name) == 0 {
		utils.Error("Municipality name cannot be empty")
	}
	return Municipality{Id: id, Name: name}
}
