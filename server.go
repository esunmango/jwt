package main

import (
	"jwt/routers"
	"jwt/settings"
	"github.com/urfave/negroni"
	"net/http"
)

func main(){
	settings.Init()
	router := routers.InitRouters()
	n:= negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000",n)
}
