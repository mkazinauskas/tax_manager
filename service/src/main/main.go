package main

import (
	"main/tax_manager/api"
	"main/tax_manager/datasource"
	"main/tax_manager/data"
)

func init(){
	datasource.Database{}.CheckConnection()
	data.InitDefaultData()
}

func main() {
	api.Initialize()
}
