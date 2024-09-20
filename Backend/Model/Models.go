package model

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	SecretKey = "tYrEQins27rw0ehqkKfJE0Ofxyd6r8QISFtpomcIILFUfRacmDuBa3nS9NXTpZfV99E1AEaU"
)

type Auth struct {
	Id              string `json:"Id"`
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
}

type Register struct {
	Auth      Auth   `json:"Auth"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	BirthDate string `json:"BirthDate"`

	// OPTIONNAL
	ProfilePicture string `json:"ProfilePicture"`
	Username       string `json:"Username"`
	AboutMe        string `json:"AboutMe"`
	Gender         string `json:"Gender"`
}

type Post struct {
	Id           string `json:"Id"`
	AuthorId     string `json:"AuthorId"`
	Text         string `json:"Text"`
	Image        string `json:"Image"`
	CreationDate string `json:"CreationDate"`
	IsGroup      string `json:"IsGroup"`
}

type Comment struct {
	Id           string `json:"Id"`
	AuthorId     string `json:"AuthorId"`
	Text         string `json:"Text"`
	CreationDate string `json:"CreationDate"`
	PostId       string `json:"PostId"`
}

type ResponseWriter struct {
	http.ResponseWriter
}

/*
This function takes 1 argument:
  - a string who contain a description of the error

The purpose of this function is to Return an error of the application who have make a request to the server.

The function return a string to the user but have no return for the server
*/
func (w *ResponseWriter) Error(err string) {
	time.Sleep(2 * time.Second)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"Error":   http.StatusText(http.StatusUnauthorized),
		"Message": err,
	})
}
