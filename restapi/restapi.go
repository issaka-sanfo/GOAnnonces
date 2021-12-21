package restapi


import (
	"github.com/gorilla/mux"
	"goanonnonces.leboncoin.fr/handler"
)


// Init the mux router
var Router = mux.NewRouter()

func Route() {

	// Route handles & endpoints
    
	// Créer une Annonce
	Router.HandleFunc("/annonces/", handler.CreateAnnonce).Methods("POST")

	// Récupérer une Annonce
	Router.HandleFunc("/annonces/{annonceid}", handler.ReadAnnonce).Methods("GET")

	// Supprimer une Annonce
	Router.HandleFunc("/annonces/{annonceid}", handler.DeleteAnnonce).Methods("DELETE")

	// Recupérer 10 Annonces
	Router.HandleFunc("/annonces/", handler.GetAnnonces).Methods("GET")

}