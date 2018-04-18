package controller

import (
	"jwt/services"
	"jwt/services/models"
	"encoding/json"
	"net/http"
	"golang.org/x/crypto/ocsp"
)

func Login(w http.ResponseWriter, r *http.Request)  {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := services.Login(requestUser)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	w.Header().Set("Content-Type", "application/json")
	w.Write(services.RefreshToken(requestUser))
}

func Logout( w http.ResponseWriter, r *http.Request, next http.HandlerFunc)  {
	err := services.Logout(r)
	w.Header().Set("Content-Type","application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
