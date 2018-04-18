package services

import (
	"jwt/api/parameters"
	"jwt/core/authentication"
	"jwt/services/models"
	"encoding/json"
	jwt  "github.com/dgrijalva/jwt-go"
	requestjwt "github.com/dgrijalva/jwt-go/request"
	"net/http"
	"text/scanner"
)

func Login(requestUser *models.User) (int, []byte){
	authBackend := authtentication.InitJWTAuthenticationBackend()

	if authBackend.Authenticate(requestUser) {
		token, er := authBackend.GenerateToken(requestUser.UUID)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _:= json.Marshal(parameters.TokenAuthentication{token})
			return http.StatusOK, response
		}
	}

	return http.StatusUnauthorized, []byte("")

}

func RefreshToken(requestUser *models.User) []byte  {
	authBackend := authentication.InitJWTAuthenticationBackend()

	token, err := authBackend.GenerateToken(requestUser.UUID)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(parameters.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

func Logout (req *http.Request) error {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := requestjwt.ParseFromRequest(req,
		func(token *jwt.Token) (interface{},error) {
			return authBackend.PublicKey, nil
		})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logouth(tokenString, tokenRequest)
}
