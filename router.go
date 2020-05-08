package main

import (
	"github.com/gorilla/mux"
)

//NewRouter dgfdgdfgdf
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/").HandlerFunc(UpdateCounter)

	return router

}
