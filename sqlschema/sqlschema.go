package sqlschema

import (
	"goanonnonces.leboncoin.fr/dbconnection"
    "goanonnonces.leboncoin.fr/messages"
)


func CreateTables(){
	messages.PrintMessage("Started ...!")

	db := dbconnection.DBsetup()
	rescreate, errcreate := db.Query(""+
		// Preparation des tables
		// Supprimer d'abord les Tables:  annonce, model, marque et categorie pour en suite ajouter des données pour les test APIs
		"DROP TABLE IF EXISTS annonce, model, marque, categorie CASCADE;"+ 
		"CREATE TABLE categorie(id INT PRIMARY KEY, libelle VARCHAR(10));"+
		"CREATE TABLE marque(id INT PRIMARY KEY, libelle VARCHAR(10));"+
		"CREATE TABLE model(id SERIAL PRIMARY KEY, libelle VARCHAR(10), "+
							"marque INT REFERENCES marque(id));"+
		"CREATE TABLE annonce(id SERIAL PRIMARY KEY, titre VARCHAR(10), contenu VARCHAR(1000), "+
							"categorie INT REFERENCES categorie(id), "+
							"marque INT REFERENCES model(id))")
	// Insertion des Categories, Marques et Models données
	db.Query("INSERT INTO categorie(id, libelle) VALUES(1, 'Emploi'),(2, 'Automobile'),(3, 'Immobilier')")
	db.Query("INSERT INTO marque(id, libelle) VALUES(1, 'Audi'),(2, 'BMW'),(3, 'Citroen');")
	db.Query("INSERT INTO model(libelle, marque) VALUES( 'Cabriolet', 1)"+
	"( 'Q2', 1),( 'Q3', 1),( 'Q5', 1),( 'Q7', 1),( 'Q8', 1),( 'R8', 1),( 'Rs3', 1),"+
	"( 'Rs4', 1),( 'Rs5', 1),( 'Rs7', 1),( 'S3', 1),( 'S4', 1),( 'S4 Avant', 1),( 'S4 Cabriolet', 1),"+
	"( 'S5', 1),( 'S7', 1),( 'S8', 1),( 'SQ5', 1),( 'SQ7', 1),( 'Tt', 1),( 'Tts', 1),( 'V8', 1)"+ // Fin Modèls Audi
	",( 'M3', 2),( 'M4', 2),( 'M5', 2),( 'M535', 2),( 'M6', 2),( 'M635', 2),( 'Serie 1', 2),( 'Serie 2', 2),( 'Serie 3', 2),( 'Serie 4', 2),"+
	"( 'Serie 5', 2),( 'Serie 6', 2),( 'Serie 7', 2),( 'Serie 8', 2)"+
	",( 'C1', 3),( 'C15', 3),( 'C2', 3),( 'C25', 3),( 'C25D', 3),( 'C25E', 3),( 'C25TD', 3),( 'C3', 3),( 'C3 Aircross', 3),( 'C3 Picasso', 3),"+
	"( 'C4 Picasso', 3),( 'C5', 3),( 'C6', 3),( 'C8', 3),( 'Ds3', 3),( 'Ds4', 3),( 'Ds5', 3);")
	// Insertion d'examples d'Annonces
	db.Query("INSERT INTO annonce(titre, contenu, categorie, marque) "+
	"VALUES('Voiture 1', 'En vente à bon prix 1', 1, 1),( 'Voiture 2', 'En vente à bon prix 2', 2, 2),( 'Voiture 3', 'En vente à bon prix 3', 3, 3)")
							

	if rescreate != nil {
		messages.PrintMessage("Tables Created!")
	}



	messages.CheckError(errcreate)

}