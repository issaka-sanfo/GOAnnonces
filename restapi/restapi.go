package restapi


import (
	"github.com/gorilla/mux"
	"goanonnonces.leboncoin.fr/handler"
)


// Init the mux router
var Router = mux.NewRouter()

// Route handles & endpoints
func Route() {
    
	// Créer une Annonce
	Router.HandleFunc("/annonces/", handler.CreateAnnonce).Methods("POST")

	// Modifier une Annonce
	Router.HandleFunc("/annonces/", handler.UpdateAnnonce).Methods("PUT")

	// Récupérer une Annonce
	Router.HandleFunc("/annonces/{annonceid}", handler.ReadAnnonce).Methods("GET")

	// Supprimer une Annonce
	Router.HandleFunc("/annonces/{annonceid}", handler.DeleteAnnonce).Methods("DELETE")

	// Recupérer 10 Annonces
	Router.HandleFunc("/annonces/", handler.GetAnnonces).Methods("GET")

}