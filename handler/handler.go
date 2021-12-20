package handler

import (
	"net/http"
	"encoding/json"
	"goanonnonces.leboncoin.fr/model"
	"goanonnonces.leboncoin.fr/dbconnection"
	"goanonnonces.leboncoin.fr/messages"
)


// Get all annonces
// response and request handlers
func GetAnnonces(w http.ResponseWriter, r *http.Request) {
    db := dbconnection.DBsetup()

    messages.PrintMessage("Getting annonces...")

    // Get all movies from movies table that don't have movieID = "1"
    rows, err := db.Query("SELECT * FROM annonce")

    // check errors
    messages.CheckError(err)

    // var response []JsonResponse
    var annonces []model.Annonce

    // Foreach movie
    for rows.Next() {
        var id int
        var annonceTitre string
        var annonceContenu string
		var annonceCategorie string
		var annonceModel string

        err = rows.Scan(&id, &annonceTitre, &annonceContenu, &annonceCategorie, &annonceModel)

        // check errors
        messages.CheckError(err)

        annonces = append(annonces, model.Annonce{Titre: annonceTitre, Contenu: annonceContenu, Categorie: annonceCategorie, Model: annonceModel})
    }

    var response = model.JsonResponse{Type: "success", Data: annonces}

    json.NewEncoder(w).Encode(response)
}