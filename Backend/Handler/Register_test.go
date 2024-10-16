package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	model "social-network/Model"
	utils "social-network/Utils"
)

func TestRegister(t *testing.T) {
	// Crée un mock de base de données (ou une vraie connexion en mémoire)
	db, err := utils.OpenDb("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Erreur lors de la création de la base de données en mémoire : %v", err)
		return
	}
	defer db.Close()

	rr, err := TryRegister(t, db, model.Register{
			Auth: model.Auth{
				Email:           "unemail@gmail.com",
				Password:        "MonMotDePasse123!",
				ConfirmPassword: "MonMotDePasse123!",
			},
			FirstName: "Jean",
			LastName:  "Dujardin",
			BirthDate: "1990-01-01",
		})
	if err != nil {
		t.Fatal(err)
		return
	}

	expected := "Register successfully"
	// Check the response body is what we expect.
	bodyValue := make(map[string]any)

	if err := json.Unmarshal(rr.Body.Bytes(), &bodyValue); err != nil {
		t.Fatalf("Erreur lors de la réception de la réponse de la requête : %v", err)
		return
	}
	if bodyValue["Success"] != true {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		return
	}
}

func TryRegister(t *testing.T, db *sql.DB, register model.Register) (*httptest.ResponseRecorder, error) {
	// Create a table for testing
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Auth (
		Id VARCHAR(36) NOT NULL UNIQUE PRIMARY KEY,
		Email VARCHAR(100) NOT NULL UNIQUE,
		Password VARCHAR(50) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS UserInfo (
		Id VARCHAR(36) NOT NULL UNIQUE REFERENCES "Auth"("Id"),
		Email VARCHAR(100) NOT NULL UNIQUE REFERENCES "Auth"("Email"),
		FirstName VARCHAR(50) NOT NULL, 
		LastName VARCHAR(50) NOT NULL,
		BirthDate VARCHAR(20) NOT NULL,
		ProfilePicture VARCHAR(100),
		Username VARCHAR(50),
		AboutMe VARCHAR(280)  
	);
`)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(register)
	if err != nil {
		return nil, err
	}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register(db))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		return nil, fmt.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)

	}

	return rr, nil
}
