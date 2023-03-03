package routes

import (
	"github.com/gorilla/mux"
	"github.com/EleisonC/vending-machine/handlers"
)


var RegisterScheduleRoutes = func(router *mux.Router) {
	router.HandleFunc("/createuser", handlers.CreateNewUserHn).Methods("POST")
	router.HandleFunc("/loginuser", handlers.LoginUserHn).Methods("POST")
}