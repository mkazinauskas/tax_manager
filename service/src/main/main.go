package main

import (
	"main/tax_manager/datasource"
	"main/tax_manager/file"
	"main/tax_manager/factory"
	"main/tax_manager/api"
	"log"
	"main/tax_manager"
)

func init() {
	log.Println("Initializing application")
	datasource.Database{}.CheckConnection()
}

func main() {
	log.Println("Populating data from file")
	file.NewPopulateDataFromFile(factory.DefaultApplicationFactory{}).Populate(tax_manager.TAX_FILE_NAME)

	api.Initialize(factory.DefaultApplicationFactory{})
}
