package model

// Data structures

type Annonce struct {
	Id string `json:"id"`
    Titre string `json:"titre"`
    Contenu string `json:"contenu"`
	Categorie string `json:"categorie"`
	Model string `json:"model"`
}

type Categorie struct {
	Libelle string `json:"libelle"`
}

type Marque struct {
	Libelle string `json:"libelle"`
}

type Model struct {
	Libelle string `json:"libelle"`
	Marque string `json:"marque"`
}

type JsonResponse struct {
    Type    string `json:"type"`
    Data    []Annonce `json:"data"`
    Message string `json:"message"`
}