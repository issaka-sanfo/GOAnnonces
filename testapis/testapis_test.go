package testapis

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"goanonnonces.leboncoin.fr/handler"
)


func TestReadAnnonceByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/annoces/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.ReadAnnonce)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"type":"succès","data":[{"id":"1","titre":"Voiture 1","contenu":"En vente à bon prix 1","categorie":"1","model":"1"}],"message":"L'Annonce a été bien recupérée!"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


func TestGetAnnonces(t *testing.T) {
	req, err := http.NewRequest("GET", "/entries", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.GetAnnonces)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"type":"success","data":[{"id":"1","titre":"Voiture 1","contenu":"En vente à bon prix 1","categorie":"1","model":"1"},{"id":"2","titre":"Voiture 2","contenu":"En vente à bon prix 2","categorie":"2","model":"2"},{"id":"3","titre":"Voiture 3","contenu":"En vente à bon prix 3","categorie":"3","model":"3"}],"message":"La liste des Annonces"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}