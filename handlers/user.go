package handlers

import (
	"github.com/afifialaa/USER-AUTH/auth"
	"github.com/afifialaa/USER-AUTH/models"
	session "github.com/afifialaa/USER-AUTH/sessions"
	"github.com/afifialaa/USER-AUTH/validation"

	"encoding/json"
	"fmt"
	"net/http"
)

// Login handle
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logging user in")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	user := models.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// Validate user
	validUser := validation.ValidateUserLogin(&user)
	if !validUser {
		// Send failed response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		data := map[string]string{"msg": "not a valid user input"}
		json.NewEncoder(w).Encode(data)
		return

	}

	_, err := user.Find()
	// User was not found
	if err != nil {
		w.Header().Set("Content-Type", "application/json")

		data := map[string]string{"msg": "user was not found"}
		json.NewEncoder(w).Encode(data)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Store email
	session.Start(user.Email)
	// Create token
	token := auth.GenerateToken(user.Email)
	data := map[string]string{"msg": "user found", "token": token}
	json.NewEncoder(w).Encode(data)
	return
}

func SignoutHandle(w http.ResponseWriter, r *http.Request) {
	session.End()
	fmt.Println("session was ended")
}

/* Create a new user */
func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("#Signup_handle")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	user := models.User{
		r.FormValue("email"),
		r.FormValue("password"),
	}

	// valid := validation.ValidateUser(&user)
	valid := true
	if valid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		_, err := user.Create()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			data := map[string]string{"err": "failed to create user"}
			json.NewEncoder(w).Encode(data)
			return
		}

		// Generate token -> not need to generate token
		data := map[string]string{"msg": "user was created successfully"}
		json.NewEncoder(w).Encode(data)
	} else {
		w.Header().Set("Content-Type", "application/json")
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
