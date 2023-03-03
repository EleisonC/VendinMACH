package helpers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"os"

	"github.com/EleisonC/vending-machine/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func VenErrorHandler(w http.ResponseWriter, errMes string, err error) {
	errRes := models.ErrMessageRes{
		Message: errMes,
		RawErrorMessage: err.Error(),
	}
	res, _ := json.Marshal(errRes)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(res)
}

func GenerateJWT(role string) (string, error) {
	signMeth := jwt.SigningMethodHS256
	secretKeyUse := "work"
	// token := jwt.New(signMeth)

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(90 * time.Minute).Unix()
	claims["username"] = role

	
	token := jwt.NewWithClaims(signMeth, claims)
	tokenString,err := token.SignedString([]byte(secretKeyUse))
	if err != nil {
		return "", err
	}

	return tokenString,nil
}

func LoadEnvVal(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}


