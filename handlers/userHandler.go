package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EleisonC/vending-machine/helpers"
	"github.com/EleisonC/vending-machine/models"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func CreateNewUserHn(w http.ResponseWriter, r *http.Request){
	var user models.UserModel
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.VenErrorHandler(w, "User Not Created", err)
		fmt.Println(err.Error())
		return
	}

	err := validate.Struct(&user)
	if err != nil {
		helpers.VenErrorHandler(w, "User Not Created", err)
		return
	}

	err = models.CreateNewUser(&user)
	if err != nil {
		helpers.VenErrorHandler(w, "Something might have happned during creation", err)
		return
	}

	postRes := models.PosMessageRes{
		Message: "User created",
	}

	res, err := json.Marshal(postRes)
	if err != nil {
		helpers.VenErrorHandler(w, "Somthing Happened. But User Is Create", err)
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func LoginUserHn(w http.ResponseWriter, r *http.Request) {
	var userLogin models.UserLogIn
	var user models.UserModeldb

	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		helpers.VenErrorHandler(w, "Failed Login Attempt", err)
	}

	err := validate.Struct(&userLogin)
	if err != nil {
		helpers.VenErrorHandler(w, "Failed Login Attempt", err)
		return
	}

	// query the user to get hashed password
	err = models.FindUserByUsername(userLogin.Username, &user)
	if err != nil {
		helpers.VenErrorHandler(w, "Failed Login Attempt", err)
		return
	}

	// compare hased password
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(userLogin.Password))
	if err != nil {
		helpers.VenErrorHandler(w, "Invalid Password Or Username", err)
		return
	}


	// GenerateJWT

	tokenSt, err := helpers.GenerateJWT(user.Role)
	if err != nil {
		helpers.VenErrorHandler(w, "Invalid Password Or Username", err)
		return
	}

	loginRes := models.TokenRes{
		Message: "User Authenticated",
		TokenString: tokenSt,
	}
	res, err := json.Marshal(loginRes)
	if err != nil {
		helpers.VenErrorHandler(w, "Issue Authenticating", err)
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}