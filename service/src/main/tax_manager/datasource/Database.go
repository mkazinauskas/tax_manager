package datasource

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"main/tax_manager/utils"
	"fmt"
)

func CheckConnection() {
	db := acquireConnection()
	err := db.Ping()
	if err != nil {
		panic(err.Error())
	}
	closeConnection(db)
}

func Execute(query string, args ... interface{}) (sql.Result) {
	fmt.Println(query, args)
	db := acquireConnection()
	stmt, prepareError := db.Prepare(query)
	utils.Check(prepareError)

	result, executionError := stmt.Exec(args...)
	utils.Check(executionError)

	return result
}

func Query(query string, args ... interface{}) (*sql.Rows) {
	fmt.Println(query, args)
	db := acquireConnection()
	stmt, prepareError := db.Prepare(query)
	utils.Check(prepareError)

	result, executionError := stmt.Query(args...)
	utils.Check(executionError)

	return result
}

func acquireConnection() *sql.DB {
	db, fault := sql.Open("mysql", "root:@/tax_manager")
	utils.Check(fault)
	return db
}

func closeConnection(db *sql.DB) {
	db.Close()
}


