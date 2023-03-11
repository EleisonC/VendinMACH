package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EleisonC/vending-machine/helpers"
	"github.com/EleisonC/vending-machine/models"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/mux"
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
	tokenSt, err := helpers.GenerateJWT(user.Role, user.Username, user.Id)
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

func EditUserDataHn(w http.ResponseWriter, r *http.Request) {
	var editUser models.EditUser
	params := mux.Vars(r)
	userId := params["userId"]
	if err := json.NewDecoder(r.Body).Decode(&editUser); err != nil {
		helpers.VenErrorHandler(w, "User Not Updated", err)
		return
	}

	if err := validate.Struct(&editUser); err != nil {
		helpers.VenErrorHandler(w, "User Not Updated Step 2", err)
		return
	}

	// verify claims
	usernameST, err := helpers.ExtractClaims(w, r)
	if err != nil {
		helpers.VenErrorHandler(w, "Claims Issue", err)
		return
	}

	if usernameST["userId"] != userId {
		errMessage := models.PosMessageRes{
			Message: "Not Authorized",
		}
		res, _ := json.Marshal(errMessage)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write(res)
		return
	}
	
	if err := models.UpdateUser(&editUser, userId); err != nil {
		helpers.VenErrorHandler(w, "Something Happened During Update", err)
		return
	}

	postRes := models.PosMessageRes{
		Message: "User created",
	}

	res, err := json.Marshal(postRes)
	if err != nil {
		helpers.VenErrorHandler(w, "Somthing Happened. But User Is Updated", err)
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserByUserNameHn(w http.ResponseWriter, r *http.Request) {
	var user models.UserModeldb
	type Username struct {
		Username string `json:"username"`
	}
	var username Username
	if err := json.NewDecoder(r.Body).Decode(&username); err != nil {
		helpers.VenErrorHandler(w, "User Not Found", err)
		fmt.Println(err.Error())
		return
	}

	err := models.FindUserByUsername(username.Username, &user)
	if err != nil {
		helpers.VenErrorHandler(w, "Issue Getting User", err)
		return
	}

	// extract claims
	usernameST, err := helpers.ExtractClaims(w, r)
	if err != nil {
		helpers.VenErrorHandler(w, "Claims Issue", err)
		return
	}

	if usernameST["username"] != user.Username {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		helpers.VenErrorHandler(w, "Issue Getting User", err)
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ChangePasswordHn(w http.ResponseWriter, r *http.Request) {
	var passChange models.PasswordChange
	var user models.UserModeldb
	params := mux.Vars(r)
	userId := params["userId"]

	if err := json.NewDecoder(r.Body).Decode(&passChange); err != nil {
		helpers.VenErrorHandler(w, "Failed to change password", err)
		return
	}

	if err := validate.Struct(&passChange); err != nil {
		helpers.VenErrorHandler(w, "The passwords are faulty", err)
		return
	}

	err := models.FindUserById(userId, &user)
	if err != nil {
		helpers.VenErrorHandler(w, "Failed Login Attempt", err)
		return
	}

	//verify claims
	usernameST, err := helpers.ExtractClaims(w, r)
	if err != nil {
		helpers.VenErrorHandler(w, "Claims Issue", err)
		return
	}

	if usernameST["username"] != user.Username {
		errMessage := models.PosMessageRes{
			Message: "Not Authorized",
		}
		res, _ := json.Marshal(errMessage)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write(res)
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(passChange.OldPassword))
	if err != nil {
		helpers.VenErrorHandler(w, "Invalid Password", err)
		return
	}


	err = models.UpdateUserPass(userId, passChange.NewPassword)
	if err != nil {
		helpers.VenErrorHandler(w, "Invalid Password", err)
		return
	}

	postRes := models.PosMessageRes{
		Message: "User pass changed",
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

func DepositCoinsHn(w http.ResponseWriter, r *http.Request) {
	var depositValue *models.Deposit
	var user models.UserModeldb
	if err := json.NewDecoder(r.Body).Decode(&depositValue); err != nil {
		helpers.VenErrorHandler(w, "Not Accepted Coin It should [20, 10, 5, 50, 100]", err)
	}

	if err := validate.Struct(&depositValue); err != nil {
		helpers.VenErrorHandler(w, "Not Accepted Coin It should [20, 10, 5, 50, 100]", err)
	}


	// Extract and verify claims
	usernameST, err := helpers.ExtractClaims(w, r)
	if err != nil {
		helpers.VenErrorHandler(w, "Claims Issue", err)
		return
	}
	userId := usernameST["userId"].(string)

	err = models.FindUserById(userId, &user)
	if err != nil {
		helpers.VenErrorHandler(w, "Failed Login Attempt", err)
		return
	}

	// new deposit value oldValue + depositValue.Deposit
	newDepositValue := depositValue.Deposit + user.Deposit

	err = models.DepositCoin(userId, newDepositValue)
	if err != nil {
		helpers.VenErrorHandler(w, "Failed to deposit cash", err)
		return
	}

	postRes := models.PosMessageRes{
		Message: "Deposit Made",
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

func ResetDepositHn(w http.ResponseWriter, r *http.Request) {
	// Extract and verify claims
	usernameST, err := helpers.ExtractClaims(w, r)
	if err != nil {
		helpers.VenErrorHandler(w, "Claims Issue", err)
		return
	}
	userId := usernameST["userId"].(string)
	if userRole := usernameST["role"].(string); userRole != "buyer" {
		postRes := models.PosMessageRes{
			Message: "Not Enough Rights To Make This Request",
		}

		res, err := json.Marshal(postRes)
		if err != nil {
			helpers.VenErrorHandler(w, "Somthing Happened. But User Is Create", err)
			return
		}

		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write(res)

		return
	}


	newDepositValue := 0
	err = models.DepositCoin(userId, newDepositValue)
	if err != nil {
		helpers.VenErrorHandler(w, "Failed to deposit cash", err)
		return
	}

	postRes := models.PosMessageRes{
		Message: "Deposit Reset",
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
