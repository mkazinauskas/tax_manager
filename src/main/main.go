package main

import (
	"main/tax_manager/api"
	"main/tax_manager/data"
)

func main() {
	data.Do()

	api.Initialize()

}
