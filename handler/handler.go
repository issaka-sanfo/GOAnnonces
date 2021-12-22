package handler

import (
	"net/http"
	"encoding/json"
    "fmt"
    "strings"
	"goanonnonces.leboncoin.fr/model"
	"goanonnonces.leboncoin.fr/dbconnection"
	"goanonnonces.leboncoin.fr/messages"
    "golang.org/x/text/runes"
    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
    "unicode"
    "github.com/gorilla/mux"
    "strconv"
)




//************************************ L' Algorithme de Matching ************************************//
func MatchingAlgo(clientsaisie string) int {
    db := dbconnection.DBsetup()

    id := 0
    models, err := db.Query("SELECT id, libelle, marque FROM model")
    
    // check errors
    messages.CheckError(err)

    // Pour chaque Model prédéfini
    for models.Next() {
        var modelid int
        var modellibelle string
        var modelmarque int

        err = models.Scan(&modelid, &modellibelle, &modelmarque)

        // check errors
        messages.CheckError(err)

        // Des variables Patterns pour le Model de vehicule
        temporaire := modellibelle // Récupérer le model prédéfini dans une variable temporaire
        temporaire = strings.ToLower(temporaire) // trasformer en minuscule
        patternremove := strings.Replace(temporaire, " ", "", -1) // Supprimer les espaces


        /* Verifier si Le Model saisi par l'utilisateur contient "patternremove"*/

        modeltemp := clientsaisie// Récupérer le modèle saisi par l'utilisateur dans une variable temporaire
        t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC) //Formaatage Supprimer les accents
        saisie, _, _ := transform.String(t, modeltemp) //Supprimer les accents
        saisie = strings.ToLower(saisie) // trasformer en minuscule
        saisieremove := strings.Replace(saisie, " ", "", -1) // Supprimer les espaces
        /* D'après les examples donnés; le Client saisie plus caractères pour le modèle de Vehicule que le nombre de caractères du Modèle prédéfini
        /* Alors, verifier si on trouvera le Modèle prédéfini avec le patternremove dans le Modèle saisi par l'utilisateur, si oui: On est dans le bon Modèle*/
        if strings.Contains(saisieremove, patternremove) {
            // Trouver la marque contenant le model
            brand, error := db.Query("SELECT id, libelle FROM marque WHERE id=$1 LIMIT 1", modelmarque)
            messages.CheckError(error)

            if brand.Next() {
    
                var libelle string
                err := brand.Scan(&id, &libelle)
                messages.CheckError(err)
                
            }
             
        }
    }
    return id
}



// Créer une petite Annonce
// response and request handler
func CreateAnnonce(w http.ResponseWriter, r *http.Request) {
    db := dbconnection.DBsetup()

    annonceTitre := r.FormValue("titre")
    annonceContenu := r.FormValue("contenu")
    annonceCategorie := r.FormValue("categorie")
    annonceModel := r.FormValue("model")

    // var response JsonResponse
    var response = model.JsonResponse{}
    // var response []JsonResponse
    var annonces []model.Annonce

    if annonceTitre == "" || annonceContenu == "" || annonceCategorie == "" {
        response = model.JsonResponse{Type: "erreur", Message: "Merci de remplir tous les champs titre, contenu et categorie!."}
    } else {
        if annonceCategorie == "Automobile" && annonceModel == "" {
            response = model.JsonResponse{Type: "erreur", Message: "Une Annonce avec la categorie Automobile doit avoir un model."}
        }
        if annonceCategorie != "Automobile" && annonceModel != "" {
            response = model.JsonResponse{Type: "erreur", Message: "Une Annonce avec une categorie differente de celle Automobile ne doit pas avoir de Model."}
        }

        // Verifier si la categorie saisie existe
        category, error := db.Query("SELECT id, libelle FROM categorie WHERE libelle=$1 LIMIT 1", annonceCategorie)
        messages.CheckError(error)
        var categoryid int
        var categorylibelle string
        if category.Next(){
            err := category.Scan(&categoryid, &categorylibelle)
            messages.CheckError(err)

            // Annonce Automobile
            if annonceCategorie == "Automobile" && annonceModel != "" {

                brandid := MatchingAlgo(annonceModel)
                if brandid != 0 {
                    messages.PrintMessage("Inserting Annonce into DB")
                    fmt.Println("Inserting new Annonce with Titre: " + annonceTitre + " , Contenu: " + annonceContenu + " , Categorie:" + annonceCategorie + " and Model:"+annonceModel)
                    var lastInsertID int
                    // Inserer une annonce en associant la marque trouvée
                    erreur := db.QueryRow("INSERT INTO annonce(titre, contenu, categorie, marque) VALUES($1, $2, $3, $4) returning id;", annonceTitre, annonceContenu, categoryid, brandid).Scan(&lastInsertID)
                    messages.CheckError(erreur)
                    annonces = append(annonces, model.Annonce{Titre: annonceTitre, Contenu: annonceContenu, Categorie: annonceCategorie, Model: annonceModel})
                    response = model.JsonResponse{Type: "succès", Data:annonces,  Message: "L' Annonce " +annonceCategorie+" a été bien insérée!"} 
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
                annonces = append(annonces, model.Annonce{Titre: annonceTitre, Contenu: annonceContenu, Categorie: annonceCategorie, Model: annonceModel})
                response = model.JsonResponse{Type: "succès", Data: annonces, Message: "L' Annonce " +annonceCategorie+" a été bien insérée!"}
            }

        } else {
            response = model.JsonResponse{Type: "erreur", Message: "La Categorie saisie n'existe pas! Merci de revoir."}
        }

    }

    json.NewEncoder(w).Encode(response)
}




// Modifier une petite Annonce

// response and request handler
func UpdateAnnonce(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r) // Query variable

    annonceId := params["annonceid"]
    newTitre := r.FormValue("titre")
    newContenu := r.FormValue("contenu")
    newcategorie := r.FormValue("categorie")
    newmodel := r.FormValue("model")

    var response = model.JsonResponse{}

    // var response []JsonResponse
    var annonces []model.Annonce

    if annonceId == "" {
        response = model.JsonResponse{Type: "erreur", Message: "Ajouter l'identifiant de l'Annonce dans le paramètre."}
    } else {
        db := dbconnection.DBsetup()

        row, _ := db.Query("SELECT id, titre, contenu, categorie, marque FROM annonce WHERE id = $1", annonceId)
        if row.Next() {

            var id string
            var annonceTitre string
            var annonceContenu string
            var annonceCategorie string
            var annonceModel string

            messages.PrintMessage("Reading Annonce from DB")

            err := row.Scan(&id, &annonceTitre, &annonceContenu, &annonceCategorie, &annonceModel)
            messages.CheckError(err)

             // Verifier si la categorie saisie existe
            category, error := db.Query("SELECT id, libelle FROM categorie WHERE libelle=$1 LIMIT 1", newcategorie)
            messages.CheckError(error)
            var categoryid int
            var categorylibelle string
            if category.Next(){
                err := category.Scan(&categoryid, &categorylibelle)
                messages.CheckError(err)

                // Annonce Auto
                if categorylibelle == "Automobile" && annonceModel != "" {

                    brandid := MatchingAlgo(newmodel)
                    if brandid != 0 {
                        messages.PrintMessage("Updating Annonce into DB")
                        fmt.Println("Updating new Annonce with Titre: " + annonceTitre + " , Contenu: " + annonceContenu + " , Categorie:" + annonceCategorie + " and Model:"+annonceModel)
                        var lastInsertID int
                        // Modifier une annonce en associant la marque trouvée
                        db.QueryRow("UPDATE annonce SET titre= $1, contenu= $2, categorie= $3, marque= $4 WHERE id= $5", newTitre, newContenu, categoryid, brandid, annonceId).Scan(&lastInsertID)
                        annonces = append(annonces, model.Annonce{Titre: newTitre, Contenu: newContenu, Categorie: strconv.Itoa(categoryid), Model: strconv.Itoa(brandid)})
                        response = model.JsonResponse{Type: "succès", Data:annonces,  Message: "L' Annonce " +annonceCategorie+" a été bien modifiée!"}
                    } else {

                        response = model.JsonResponse{Type: "erreur", Message: "Le Model saisi n'existe pas! Merci de revoir."}

                    }
    
                } 

                // Annonce qui n'est pas Automobile
                if annonceCategorie != "Automobile" && annonceModel == "" {
                    messages.PrintMessage("Updating Annonce into DB")
                    fmt.Println("Updating new Annonce with Titre: " + annonceTitre + " , Contenu: " + annonceContenu + " , Categorie:" + annonceCategorie)
                    var lastInsertID int
                    db.QueryRow("UPDATE annonce SET titre= $1, contenu= $2, categorie= $3 WHERE id= $4", newTitre, newContenu, categoryid, annonceId).Scan(&lastInsertID)
                    annonces = append(annonces, model.Annonce{Titre: newTitre, Contenu: newContenu, Categorie: strconv.Itoa(categoryid)})
                    response = model.JsonResponse{Type: "succès", Data: annonces, Message: "L' Annonce " +annonceCategorie+" a été bien modifiée!"}
                }

            }else {
                response = model.JsonResponse{Type: "erreur", Message: "La Nouvelle Categorie saisie n'existe pas! Merci de revoir."}
            }

        } else {

            response = model.JsonResponse{Type: "Erreur", Message: "L'identifiant de l'Annonce n'existe pas, merci de revoir!"}

        }
        
    }

    json.NewEncoder(w).Encode(response)
}



// Recupérer une Annonce
// response and request handler
func ReadAnnonce(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // Query variable

    annonceId := params["annonceid"]

    var response = model.JsonResponse{}

    // var response []JsonResponse
    var annonces []model.Annonce

    if annonceId == "" {
        response = model.JsonResponse{Type: "erreur", Message: "Ajouter l'identifiant de l'Annonce dans le paramètre."}
    } else {
        db := dbconnection.DBsetup()

        row, _ := db.Query("SELECT id, titre, contenu, categorie, marque FROM annonce WHERE id = $1", annonceId)
        if row.Next() {

            var id string
            var annonceTitre string
            var annonceContenu string
            var annonceCategorie string
            var annonceModel string

            messages.PrintMessage("Reading Annonce from DB")

            err := row.Scan(&id, &annonceTitre, &annonceContenu, &annonceCategorie, &annonceModel)

            // check errors
            messages.CheckError(err)
    
            annonces = append(annonces, model.Annonce{Id: id, Titre: annonceTitre, Contenu: annonceContenu, Categorie: annonceCategorie, Model: annonceModel})
    
            response = model.JsonResponse{Type: "succès", Data: annonces, Message: "L'Annonce a été bien recupérée!"}

        } else {

            response = model.JsonResponse{Type: "Erreur", Message: "L'identifiant de l'Annonce n'existe pas, merci de revoir!"}

        }
        
    }

    json.NewEncoder(w).Encode(response)
}




// Supprimer une Annonce
// response and request handlers
func DeleteAnnonce(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // Query variable

    annonceId := params["annonceid"]

    var response = model.JsonResponse{}

    if annonceId == "" {
        response = model.JsonResponse{Type: "erreur", Message: "Ajouter l'identifiant de l'Annonce dans le paramètre."}
    } else {
        db := dbconnection.DBsetup()

        row, _ := db.Query("SELECT id, titre, contenu, categorie, marque FROM annonce WHERE id = $1", annonceId)
        if row.Next() {
            messages.PrintMessage("Deleting Annonce from DB")

            _, err := db.Exec("DELETE FROM annonce WHERE id = $1", annonceId)

            // check errors
            messages.CheckError(err)

            response = model.JsonResponse{Type: "succès", Message: "L'Annonce a été bien supprimé!"}
        } else {
            response = model.JsonResponse{Type: "Erreur", Message: "L'identifiant de l'Annonce n'existe pas, merci de revoir!"}
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
        var id string
        var annonceTitre string
        var annonceContenu string
		var annonceCategorie string
		var annonceModel string

        err = rows.Scan(&id, &annonceTitre, &annonceContenu, &annonceCategorie, &annonceModel)

        // check errors
        messages.CheckError(err)

        annonces = append(annonces, model.Annonce{Id: id, Titre: annonceTitre, Contenu: annonceContenu, Categorie: annonceCategorie, Model: annonceModel})
    }

    var response = model.JsonResponse{Type: "succès", Data: annonces, Message: "La liste des Annonces"}

    json.NewEncoder(w).Encode(response)
}