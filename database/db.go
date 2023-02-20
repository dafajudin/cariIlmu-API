package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"cariIlmu-API/config"
)

var db *sql.DB
var err error


func Init(){
	fmt.Println("DB_NAME: ", config.GetConfig("DB_NAME"))

	connectionString := config.GetConfig("DB_USER") + ":" + config.GetConfig("DB_PASS") + "@tcp(" + config.GetConfig("DB_HOST") + ":" + config.GetConfig("DB_PORT") + ")/" + config.GetConfig("DB_NAME")

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connectionString error..")
	}

	err = db.Ping()
	if err != nil {
		panic("Database Name Invalid")
	}
}

func CreateConnection() *sql.DB {
	return db
}
