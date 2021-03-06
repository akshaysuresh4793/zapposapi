package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func connect(database string) error {
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPassword := "password"
	if len(os.Getenv("DBHOST")) > 0 {
		dbHost = os.Getenv("DBHOST")
	}
	if len(os.Getenv("DBPORT")) > 0 {
		dbPort = os.Getenv("DBPORT")
	}
	if len(os.Getenv("DBUSER")) > 0 {
		dbUser = os.Getenv("DBUSER")
	}
	if len(os.Getenv("DBPASSWORD")) > 0 {
		dbPassword = os.Getenv("DBPASSWORD")
	}
	fmt.Println("Connecting to MySQL - ", dbHost, ":", dbPort)
	var err error
	db, err = sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+database)
	handleError(err)
	cacheInitialize()
	return err
}
