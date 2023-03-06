package routes

import (
	"github.com/gorilla/mux"
	"github.com/EleisonC/vending-machine/handlers"
	"github.com/EleisonC/vending-machine/helpers"
)


var RegisterScheduleRoutes = func(router *mux.Router) {
	router.HandleFunc("/createuser", handlers.CreateNewUserHn).Methods("POST")
	router.HandleFunc("/loginuser", handlers.LoginUserHn).Methods("POST")
	router.HandleFunc("/edituser/{userId}", helpers.VerifyJWT(handlers.EditUserDataHn)).Methods("PUT")
	router.HandleFunc("/getuser", helpers.VerifyJWT(handlers.GetUserByUserNameHn)).Methods("POST")
	router.HandleFunc("/changeuserpass/{userId}", helpers.VerifyJWT(handlers.ChangePasswordHn)).Methods("PUT")
}

