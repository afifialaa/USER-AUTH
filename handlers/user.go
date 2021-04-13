package handlers

import (
	"github.com/afifialaa/user-auth/auth"
	"github.com/afifialaa/user-auth/models"
	session "github.com/afifialaa/user-auth/sessions"
	"github.com/afifialaa/user-auth/validation"

	"encoding/json"
	"fmt"
	"net/http"
)

// Login handle
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	user := models.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	validUser := validation.ValidateUserLogin(&user)
	if !validUser {
		w.WriteHeader(http.StatusUnauthorized)
		data := map[string]string{"msg": "Unauthorized user"}
		json.NewEncoder(w).Encode(data)
		return

	}

	_, err := user.Find()
	if err != nil {

		w.WriteHeader(http.StatusNotFound)
		data := map[string]string{"msg": "User was not found"}
		json.NewEncoder(w).Encode(data)
		return
	}

	// session.Start(user.Email)

	// Generate token
	token := auth.GenerateToken(user.Email)

	w.WriteHeader(http.StatusOK)
	data := map[string]string{"token": token}
	json.NewEncoder(w).Encode(data)
	return
}

func SignoutHandle(w http.ResponseWriter, r *http.Request) {
	session.End()
	fmt.Println("session was ended")
}

/* Create a new user */
func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user := models.User{
		r.FormValue("email"),
		r.FormValue("password"),
	}

	// valid := validation.ValidateUser(&user)
	valid := true
	w.Header().Set("Content-Type", "application/json")
	if valid {

		_, err := user.Create()
		if err != nil {

			// Email already in use
			w.WriteHeader(http.StatusBadRequest)
			data := map[string]string{"err": "Failed to create user"}
			json.NewEncoder(w).Encode(data)
			return
		}

		// Valid user
		w.WriteHeader(http.StatusCreated)
		data := map[string]string{"msg": "User was created successfully"}
		json.NewEncoder(w).Encode(data)
	} else {

		// Not a valid input
		w.WriteHeader(http.StatusBadRequest)
		data := map[string]string{"err": "Not a valid user"}
		json.NewEncoder(w).Encode(data)
		return
	}
}

/* Test endpoint */
func TestHandle(w http.ResponseWriter, r *http.Request) {

	var token string = auth.GetToken(r)
	// No token found
	if token == "" {
		w.Header().Set("Content-Type", "application/json")
		data := map[string]string{"err": "token was not found"}
		json.NewEncoder(w).Encode(data)
		return
	}

	// Validate token
	validToken := auth.VerifyToken(token)

	// Not a valid token
	if !validToken {
		w.Header().Set("Content-Type", "application/json")
		data := map[string]string{"err": "invalid token"}
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
