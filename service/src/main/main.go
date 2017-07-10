package main

import (
	"main/tax_manager/api"
	"main/tax_manager/datasource"
	"main/tax_manager/data"
)

func init(){
	datasource.CheckConnection()
	data.InitDefaultData()
}

func main() {
	api.Initialize()
}
