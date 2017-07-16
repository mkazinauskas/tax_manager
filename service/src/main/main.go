package main

import (
	"main/tax_manager/datasource"
	"main/tax_manager/file"
	"main/tax_manager/factory"
	"main/tax_manager/api"
)

func init() {
	datasource.Database{}.CheckConnection()
}

func main() {
	file.NewPopulateDataFromFile(factory.DefaultApplicationFactory{}).Populate("tax_file.csv")
	api.Initialize(factory.DefaultApplicationFactory{})
}
