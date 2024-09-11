package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "social-network/Model"
	utils "social-network/Utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	contextValue := r.Context().Value(model.RegisterCtx).([]byte)

	var register model.Register
	if err := json.Unmarshal(contextValue, &register); err != nil {
		fmt.Println(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Internal Error: There is an Unmarshal error")
		return
	}

	// INSERT in BDD in the Auth table and in the UserInfo table
	if err := utils.InsertIntoDb("Auth", register.Auth.Id, register.Auth.Email, register.Auth.Password); err != nil {
		fmt.Println(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Internal Error: There is a probleme during the push in the DB")
		return
	}

	if err := utils.InsertIntoDb("UserInfo", register.Auth.Id, register.Auth.Email, register.FirstName, register.LastName, register.BirthDate, register.ProfilePicture, register.Username, register.AboutMe); err != nil {
		fmt.Println(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Internal Error: There is a probleme during the push in the DB")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(register)
}
