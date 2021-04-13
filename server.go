package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/afifialaa/user-auth/config"
	database "github.com/afifialaa/user-auth/database"
	"github.com/afifialaa/user-auth/handlers"
)

type Status struct {
	msg string
}

func main() {

	config.SetEnv()

	// routes
	http.HandleFunc("api/user/signup", handlers.Signup)
	http.HandleFunc("api/user/login", handlers.Login)
	database.Connect()

	// listening for requests
	fmt.Println("server is running")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
