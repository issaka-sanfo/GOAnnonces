package handler

import (
	"net/http"
	"encoding/json"
    "fmt"
    "strings"
	"goanonnonces.leboncoin.fr/model"
	"goanonnonces.leboncoin.fr/dbconnection"
	"goanonnonces.leboncoin.fr/messages"
    "github.com/asaskevich/govalidator"
    "golang.org/x/text/runes"
    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
    "unicode"
    "github.com/gorilla/mux"
)

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

                //************************************ L' Algorithme de Matching ************************************//

                // Selectionner tous les modèles de vehicule
                // Requete
                rows, err := db.Query("SELECT id, libelle, marque FROM model")

                // check errors
                messages.CheckError(err)

                // Pour chaque Model prédéfini
                for rows.Next() {
                    var modelid int
                    var modellibelle string
                    var modelmarque string

                    err = rows.Scan(&modelid, &modellibelle, &modelmarque)

                    // check errors
                    messages.CheckError(err)

                    // Des variables Patterns pour le Model de vehicule
                    temporaire := modellibelle // Récupérer le model prédéfini dans une variable temporaire
                    temporaire = strings.ToLower(temporaire) // trasformer en minuscule
                    pattern1remove := strings.Replace(temporaire, " ", "", -1)
                    pattern2add := ""
                    for i := 0; i < len(temporaire); i++ { // Parcourir les caractères du model et si on croise un entier on le precède avec de l'espace
                        if govalidator.IsInt(string(temporaire[i])) { // Le caractère est un entier
                            pattern2add += " "+string(temporaire[i:len(temporaire)]) // Ajouter un espace et le reste des chiffres
                            break
                        } else {
                            pattern2add += string(temporaire[i]) // Le caractère n'est pas un entier
                        }
                    }

                    // Verifier si Le Model saisi par l'utilisateur contient "Le libelle du Model" ou un des "2 Pattens: pattern1remove et pattern2add"
                    modeltemp := annonceModel// Récupérer le modèle saisi par l'utilisateur dans une variable temporaire
                    t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
                    saisie, _, _ := transform.String(t, modeltemp)
                    saisie = strings.ToLower(saisie) // trasformer en minuscule
                    /* D'après les examples donnés; le Client saisie plus caractères pour le modèle de Vehicule que le nombre de caractères du Modèle prédéfini
                    /* Alors, verifier si on trouvera le Modèle prédéfini avec tous ses patterns dans le Modèle saisi par l'utilisateur, si oui: On est dans le bon Modèle*/
                    if strings.Contains(saisie, temporaire) || strings.Contains(saisie, pattern1remove) || strings.Contains(saisie, pattern2add)  {
                    
                        // Trouver la marque contenant le model
                        brand, error := db.Query("SELECT id, libelle FROM marque WHERE id=$1 LIMIT 1", modelmarque)
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
                            annonces = append(annonces, model.Annonce{Titre: annonceTitre, Contenu: annonceContenu, Categorie: annonceCategorie, Model: annonceModel})
                            response = model.JsonResponse{Type: "success", Data:annonces,  Message: "L' Annonce " +annonceCategorie+" a été bien insérée!"}
                        }

                        break // Arrêter au bon Matching de modèle
                        
                    } else {

                        response = model.JsonResponse{Type: "erreur", Message: "Le Model saisi n'existe pas! Merci de revoir."}

                    }
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
                response = model.JsonResponse{Type: "success", Data: annonces, Message: "L' Annonce " +annonceCategorie+" a été bien insérée!"}
            }

        } else {
            response = model.JsonResponse{Type: "erreur", Message: "La Categorie saisie n'existe pas! Merci de revoir."}
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

    var response = model.JsonResponse{Type: "success", Data: annonces, Message: "La liste des Annonces"}

    json.NewEncoder(w).Encode(response)
}