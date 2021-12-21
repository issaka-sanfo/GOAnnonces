package testapis

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"goanonnonces.leboncoin.fr/handler"
)

// "rs4 avant", l'annonce devra afficher "Audi" et "Rs4"
func TestMatchingAlgo1(t *testing.T){
	saisie := "rs4 avant"
	idbrand := handler.MatchingAlgo(saisie)

	// La marque Audi a l'identifiant "1"
	expected := 1

	if idbrand != expected {
		t.Errorf("handler returned wrong status code: got %v want %v",
		idbrand, expected)
	}
}

// "Gran Turismo Série5", l'annonce devra afficher "BMW" et "Serie 5"
func TestMatchingAlgo2(t *testing.T){
	saisie := "Gran Turismo Série5"
	idbrand := handler.MatchingAlgo(saisie)

	// La marque BMW a l'identifiant "2"
	expected := 2

	if idbrand != expected {
		t.Errorf("handler returned wrong status code: got %v want %v",
		idbrand, expected)
	}
}

// "ds 3 crossback", l'annonce devra afficher "Citroen" et "Ds3"
func TestMatchingAlgo3(t *testing.T){
	saisie := "ds 3 crossback"
	idbrand := handler.MatchingAlgo(saisie)

	// La marque Citroen a l'identifiant "3"
	expected := 3

	if idbrand != expected {
		t.Errorf("handler returned wrong status code: got %v want %v",
		idbrand, expected)
	}
}

// "CrossBack ds 3", l'annonce devra afficher "Citroen" et "Ds3"
func TestMatchingAlgo4(t *testing.T){
	saisie := "CrossBack ds 3"
	idbrand := handler.MatchingAlgo(saisie)

	// La marque Citroen a l'identifiant "3"
	expected := 3

	if idbrand != expected {
		t.Errorf("handler returned wrong status code: got %v want %v",
		idbrand, expected)
	}
}

func TestGetAnnonces(t *testing.T) {
	req, err := http.NewRequest("GET", "/annonces", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.GetAnnonces)
	handler.ServeHTTP(rr, req)
	// Verifier si la reponse est OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got = %v want = %v",
			status, http.StatusOK)
	}
}

func TestReadAnnonceByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/annonces/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.ReadAnnonce)
	handler.ServeHTTP(rr, req)
	// Verifier si la reponse est OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}


func TestDeleteByID(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/annonces/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.ReadAnnonce)
	handler.ServeHTTP(rr, req)
	// Verifier si la reponse est OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}