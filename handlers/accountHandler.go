package handlers

import (
	"github.com/afifialaa/USER-AUTH/auth"
	database "github.com/afifialaa/USER-AUTH/database"
	session "github.com/afifialaa/USER-AUTH/sessions"
	"github.com/afifialaa/USER-AUTH/validation"

	"encoding/json"
	"fmt"
	"net/http"
)

type user_type struct {
	firstName string
	lastName  string
	email     string
	password  string
}

// Login handle
func LoginHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	user := validation.User_login_type{
		r.FormValue("email"),
		r.FormValue("password"),
	}

	validUser := validation.ValidateUserLogin(&user)

	if !validUser {

		// Send failed response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		data := map[string]string{"msg": "not a valid user input"}
		json.NewEncoder(w).Encode(data)

	} else {
		userFound := database.FindUser(&user)
		if userFound {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// Start session.
			session.Start(user.Email)

			// Create token
			token := auth.GenerateToken(user.Email)

			// Generate json
			data := map[string]string{"msg": "user found", "token": token}

			// Sending response
			json.NewEncoder(w).Encode(data)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// Generate json
			data := map[string]string{"msg": "user was not found"}

			// Sending response
			json.NewEncoder(w).Encode(data)
		}

	}
}

func SignoutHandle(w http.ResponseWriter, r *http.Request) {
	session.End()
	fmt.Println("session was ended")
}

func SignupHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("#Signup_handle")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	user := validation.User_type{
		r.FormValue("firstName"),
		r.FormValue("lastName"),
		r.FormValue("email"),
		r.FormValue("password"),
	}

	valid := validation.ValidateUser(&user)
	if valid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		saved := database.SaveUser(&user)
		if !saved {
			w.Header().Set("Content-Type", "application/json")
			data := map[string]string{"msg": "user was not created"}
			json.NewEncoder(w).Encode(data)
			return
		}

		// Generate token
		token := auth.GenerateToken(user.Email)
		// Generate json
		data := map[string]string{"msg": "saved user", "token": token}

		// Sending response
		json.NewEncoder(w).Encode(data)
	} else {
		w.Header().Set("Content-Type", "application/json")

		data := map[string]string{"msg": "user was not created"}
		json.NewEncoder(w).Encode(data)
	}
}

func TestHandle(w http.ResponseWriter, r *http.Request) {

	var token string = auth.GetToken(r)
	// No token found
	if token == "" {
		w.Header().Set("Content-Type", "application/json")
		data := map[string]string{"msg": "token was not found"}
		json.NewEncoder(w).Encode(data)
		return
	}

	// Validate token
	validToken := auth.Test(token)

	// Not a valid token
	if !validToken {
		w.Header().Set("Content-Type", "application/json")
		data := map[string]string{"msg": "invalid token"}
		json.NewEncoder(w).Encode(data)
		return
	} else {
		// Serve the user
		w.Header().Set("Content-Type", "application/json")
		data := map[string]string{"msg": "token was valid and user was served"}
		json.NewEncoder(w).Encode(data)
		return
	}

}
