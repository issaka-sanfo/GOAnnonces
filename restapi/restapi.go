package restapi


import (
	"github.com/gorilla/mux"
	"goanonnonces.leboncoin.fr/handler"
)


// Init the mux router
var Router = mux.NewRouter()

func Route() {

	// Route handles & endpoints

	// Get all Annonces
	Router.HandleFunc("/annonces/", handler.GetAnnonces).Methods("GET")

}