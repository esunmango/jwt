package routers

import (
	"jwt/controller"
	"jwt/core/authentication"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router )  {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(controller.HelloController),
		)).Methods("GET")

		return router
}