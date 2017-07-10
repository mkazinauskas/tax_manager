package datasource

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func CheckConnection() {
	db := acquireConnection()
	err := db.Ping()
	if err != nil {
		panic(err.Error())
	}
	closeConnection(db)
}



func acquireConnection() *sql.DB {
	DB, err := sql.Open("mysql", "root:@/tax_manager")

	if err != nil {
		panic(err.Error())
	}
	return DB
}

func closeConnection(db *sql.DB) {
	db.Close()
}
