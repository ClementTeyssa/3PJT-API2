package main

import (
	"github.com/ClementTeyssa/3PJT-API2/controllers"
	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("POST").Path("/").Name("Index").HandlerFunc(controllers.DoVerifications)
	return router
}
