package routers

import (
	"jwt/controller"
	"jwt/core/authenticaton"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
)

func SetAuthenticationRoutes ( router *mux.Router) *mux.Router {
	router.HandlerFunc("/token-auth", controller.Login).Methods("POST")

	router.Handler("/refres-token-auth",
		negroni.New(
			negroni.HandlerFunc(controller.RefreshToken),
			)).Methods("GET")

	router.Handler("/logout",
		negroni.New(
			negroni.HandlerFunc(authenticaton.RequireTokenAuthentication),
			negroni.HandlerFunc(controller.Logout),
			)).Methods("GET")

	return router
}
