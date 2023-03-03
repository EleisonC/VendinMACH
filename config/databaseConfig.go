package config

import (
	"database/sql"
	"fmt"
	// "fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB(){
	db, err := sql.Open("sqlite3", "./vendingM.db")
	if err != nil {
		log.Fatal(err)
		fmt.Println("we get here")
	}

	DB = db
}


// Next step is to setup the models for json, the SQL structs and tables
