package datasource

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"main/tax_manager/utils"
	"log"
	"main/tax_manager"
)

type Database struct {
}

func (this Database) CheckConnection() {
	db := this.acquireConnection()
	err := db.Ping()
	if err != nil {
		panic(err.Error())
	}
	this.closeConnection(db)
}

func (this Database) Execute(query string, args ... interface{}) (sql.Result) {
	log.Println(query, args)
	db := this.acquireConnection()
	stmt, prepareError := db.Prepare(query)
	utils.Check(prepareError)

	result, executionError := stmt.Exec(args...)
	utils.Check(executionError)

	return result
}

func (this Database) Query(query string, args ... interface{}) (*sql.Rows) {
	log.Println(query, args)
	db := this.acquireConnection()
	stmt, prepareError := db.Prepare(query)
	utils.Check(prepareError)

	result, executionError := stmt.Query(args...)
	utils.Check(executionError)

	return result
}

func (Database) acquireConnection() *sql.DB {
	db, fault := sql.Open(tax_manager.DRIVER_NAME, tax_manager.DATA_SOURCE_NAME)
	utils.Check(fault)
	return db
}

func (Database) closeConnection(db *sql.DB) {
	db.Close()
}
