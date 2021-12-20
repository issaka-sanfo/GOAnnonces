package handler

import (
	"net/http"
	"encoding/json"
    "fmt"
	"goanonnonces.leboncoin.fr/model"
	"goanonnonces.leboncoin.fr/dbconnection"
	"goanonnonces.leboncoin.fr/messages"
)

// Créer une petite Annonce
// response and request handler
func CreateAnnonce(w http.ResponseWriter, r *http.Request) {
    db := dbconnection.DBsetup()

    annonceTitre := r.FormValue("titre")
    annonceContenu := r.FormValue("contenu")
    annonceCategorie := r.FormValue("categorie")
    annonceModel := r.FormValue("model")

    var response = model.JsonResponse{}

    if annonceTitre == "" || annonceContenu == "" || annonceCategorie == "" {
        response = model.JsonResponse{Type: "erreur", Message: "Merci de remplir tous les champs titre, contenu et categorie!."}
    } else {
        if annonceCategorie == "Automobile" && annonceModel == "" {
            response = model.JsonResponse{Type: "erreur", Message: "Une Annonce avec la categorie Automobile doit avoir un model."}
        }
        if annonceCategorie != "Automobile" && annonceModel != "" {
            response = model.JsonResponse{Type: "erreur", Message: "Une Annonce avec une categorie differente de celle Automobile ne doit pas avoir de Model."}
        }

        category, error := db.Query("SELECT id, libelle FROM categorie WHERE libelle=$1 LIMIT 1", annonceCategorie)
        messages.CheckError(error)
        var categoryid int
        var categorylibelle string
        if category.Next(){
            err := category.Scan(&categoryid, &categorylibelle)
            messages.CheckError(err)
        } else {
            response = model.JsonResponse{Type: "erreur", Message: "La Categorie saisie n'existe pas! Merci de revoir."}
        }
        
        // Annonce Automobile
        if annonceCategorie == "Automobile" && annonceModel != "" {
            // Verifier si le Model saisi par le Client exist dans les Models prédefinis
            usermodel, error := db.Query("SELECT id, libelle, marque FROM model WHERE libelle=$1 LIMIT 1", annonceModel)
            messages.CheckError(error)
            if usermodel.Next() {
                var modelid int
                var libelle string
                var marque string
                err := usermodel.Scan(&modelid, &libelle, &marque)
                messages.CheckError(err)

                // Trouver la marque associée au model
                brand, error := db.Query("SELECT id, libelle FROM marque WHERE id=$1 LIMIT 1", marque)
                messages.CheckError(error)

                if brand.Next() {
                    var brandid int
                    var libelle string
                    err := brand.Scan(&brandid, &libelle)
                    messages.CheckError(err)
                    messages.PrintMessage("Inserting Annonce into DB")
                    fmt.Println("Inserting new Annonce with Titre: " + annonceTitre + " , Contenu: " + annonceContenu + " , Categorie:" + annonceCategorie + " and Model:"+annonceModel)
                    var lastInsertID int
                    // Inserer une annonce en associant la marque trouve
                    erreur := db.QueryRow("INSERT INTO annonce(titre, contenu, categorie, marque) VALUES($1, $2, $3, $4) returning id;", annonceTitre, annonceContenu, categoryid, brandid).Scan(&lastInsertID)
                    messages.CheckError(erreur)
                    response = model.JsonResponse{Type: "success", Message: "L' Annonce " +annonceCategorie+" a été bien insérée!"}
                }
                
            } else {
                response = model.JsonResponse{Type: "erreur", Message: "Le Model saisi n'existe pas! Merci de revoir."}
            }
        } 

        // Annonce qui n'est pas Automobile
        if annonceCategorie != "Automobile" && annonceModel == "" {
            messages.PrintMessage("Inserting Annonce into DB")
            fmt.Println("Inserting new Annonce with Titre: " + annonceTitre + " , Contenu: " + annonceContenu + " , Categorie:" + annonceCategorie)
            var lastInsertID int
            erreur := db.QueryRow("INSERT INTO annonce(titre, contenu, categorie) VALUES($1, $2, $3) returning id;", annonceTitre, annonceContenu, categoryid).Scan(&lastInsertID)
            messages.CheckError(erreur)
            response = model.JsonResponse{Type: "success", Message: "L' Annonce " +annonceCategorie+" a été bien insérée!"}
        }
    }
    json.NewEncoder(w).Encode(response)
}

// Recupérer 10 Annonces
// response and request handler
func GetAnnonces(w http.ResponseWriter, r *http.Request) {
    db := dbconnection.DBsetup()

    messages.PrintMessage("Getting annonces...")

    // Requete
    rows, err := db.Query("SELECT id, titre, contenu, categorie, marque FROM annonce LIMIT 10")

    // check errors
    messages.CheckError(err)

    // var response []JsonResponse
    var annonces []model.Annonce

    // Pour chaque Annonce
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