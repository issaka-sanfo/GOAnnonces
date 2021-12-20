package sqlschema

import (
	"goanonnonces.leboncoin.fr/dbconnection"
    "goanonnonces.leboncoin.fr/messages"
)


func CreateTables(){
	messages.PrintMessage("Started ...!")
	db := dbconnection.DBsetup()
	messages.PrintMessage("Continue ...!")
	rescreate, errcreate := db.Query(""+
		"CREATE TABLE IF NOT EXISTS categorie(id INT PRIMARY KEY, libelle VARCHAR(10));"+
		"CREATE TABLE IF NOT EXISTS marque(id INT PRIMARY KEY, libelle VARCHAR(10));"+
		"CREATE TABLE IF NOT EXISTS model(id SERIAL PRIMARY KEY, libelle VARCHAR(10), "+
							"marque INT REFERENCES marque(id));"+
		"CREATE TABLE IF NOT EXISTS annonce(id SERIAL PRIMARY KEY, titre VARCHAR(10), contenu VARCHAR(1000), "+
							"categorie INT REFERENCES categorie(id), "+
							"model INT REFERENCES model(id))")

	db.Query("INSERT INTO categorie(id, libelle) VALUES(1, 'Emploi'),(2, 'Automobile'),(3, 'Immobilier')")
	db.Query("INSERT INTO marque(id, libelle) VALUES(1, 'Audi'),(2, 'BMW'),(3, 'Citroen');")
	db.Query("INSERT INTO model(libelle, marque) VALUES( 'Cabriolet', 1),( 'M3', 2),( 'C1', 3);")
	db.Query("INSERT INTO annonce(titre, contenu, categorie, model) "+
	"VALUES('Voiture 1', 'En vente à bon prix 1', 1, 1),( 'Voiture 2', 'En vente à bon prix 2', 2, 2),( 'Voiture 3', 'En vente à bon prix 3', 3, 3)")
							

	if rescreate != nil {
		messages.PrintMessage("Tables Created!")
	}



	messages.CheckError(errcreate)

}