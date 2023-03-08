package helpers

import (
	"encoding/json"
	"fmt"
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

func GenerateJWT(role string, username string) (string, error) {
	signMeth := jwt.SigningMethodHS256
	secretKeyUse := "work"
	// token := jwt.New(signMeth)

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(90 * time.Minute).Unix()
	claims["user"] = role
	claims["username"] = username

	
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

func VerifyJWT(endPoint func( w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token)(interface{}, error){
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					_, err := w.Write([]byte("You are Unauthorized"))
					if err != nil {
						return nil, err
					}
				}
				return []byte("work"), nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("You are Unauthorized"))
				if err != nil {
					return
				}
			}

			if token.Valid {
				endPoint(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				_, err = w.Write([]byte("You are Unauthorized"))
				if err != nil {
					return
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("You are Unauthorized"))
			if err != nil {
				return
			}
		}
	})
	
}

func ExtractClaims(w http.ResponseWriter, r *http.Request, etclaim string) (map[string]interface{}, error) {
	if r.Header["Token"] != nil {
		tokenSt := r.Header["Token"][0]
		token, err := jwt.Parse(tokenSt, func(token *jwt.Token)(interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there's an error with the signing method")
			}
			return []byte("work"), nil
		})
		if err != nil {
			return nil, err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			return claims, nil
		}
		return nil, nil
	}
	return nil, nil
}

