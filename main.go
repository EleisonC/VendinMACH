package main

import (
	"log"
	"net/http"

	"github.com/EleisonC/vending-machine/config"
	"github.com/EleisonC/vending-machine/routes"
	// "github.com/golang-migrate/migrate/v4"
	"github.com/gorilla/mux"
)

func init(){
	config.ConnectDB()
	db := config.DB
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS usertable (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT UNIQUE, password TEXT, deposit INTEGER, role TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	// m, err := migrate.New("./migrations", "./vending.db")
    // if err != nil {
    //     log.Fatal(err)
    // }

    // if err := m.Up(); err != nil && err != migrate.ErrNoChange {
    //     log.Fatal(err)
    // }

	db.Close()
}


func main(){
	r := mux.NewRouter()
	routes.RegisterScheduleRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9090", r))
}