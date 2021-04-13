package main

import (
	"fmt"
	"net/http"

	"github.com/afifialaa/user-auth/config"
	database "github.com/afifialaa/user-auth/database"
	"github.com/afifialaa/user-auth/handlers"
	gorillaHandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Status struct {
	msg string
}

func main() {

	// Set environment variables
	config.SetEnv()

	// Connect to db
	database.Connect()

	// Set routes
	r := mux.NewRouter()
	r.HandleFunc("/api/user/signup", handlers.Signup).Methods("POST")
	r.HandleFunc("/api/user/login", handlers.Login).Methods("POST")

	// listening and serve
	fmt.Println("Server is running")
	http.ListenAndServe(":8000", gorillaHandler.CORS()(r))

}
