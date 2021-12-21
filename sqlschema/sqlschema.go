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
		"DROP TABLE IF EXISTS annonce, model, marque, categorie CASCADE;"+ // Supprimer d'abord les Tables:  annonce, model, marque et categorie pour en suite ajouter des données pour les test APIs
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
	db.Query("INSERT INTO model(libelle, marque) VALUES( 'Cabriolet', 1),( 'M3', 2),( 'C1', 3);")
	// Insertion d'examples d'Annonces
	db.Query("INSERT INTO annonce(titre, contenu, categorie, marque) "+
	"VALUES('Voiture 1', 'En vente à bon prix 1', 1, 1),( 'Voiture 2', 'En vente à bon prix 2', 2, 2),( 'Voiture 3', 'En vente à bon prix 3', 3, 3)")
							

	if rescreate != nil {
		messages.PrintMessage("Tables Created!")
	}



	messages.CheckError(errcreate)

}