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
	router.HandleFunc("/depositcoin", helpers.VerifyJWT(handlers.DepositCoinsHn)).Methods("POST")
	router.HandleFunc("/createnewproduct", helpers.VerifyJWT(handlers.CreateNewProductHn)).Methods("POST")
	router.HandleFunc("/getallproducts", helpers.VerifyJWT(handlers.GetAllProductsHn)).Methods("GET")
	router.HandleFunc("/updateproduct/{productId}", helpers.VerifyJWT(handlers.UpdateProductHn)).Methods("PUT")
	router.HandleFunc("/deleteproductbyid/{productId}", helpers.VerifyJWT(handlers.DeleteProductByIdHn)).Methods("DELETE")
}

