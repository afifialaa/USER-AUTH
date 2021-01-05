package main

import (
	"fmt"
	"log"
	"net/http"

	database "github.com/afifialaa/USER-AUTH/database"
	"github.com/afifialaa/USER-AUTH/handlers"
)

type Status struct {
	msg string
}

func main() {

	// routes
	http.HandleFunc("/user/createUser", handlers.Signup)
	http.HandleFunc("/user/login", handlers.Login)
	http.HandleFunc("/api/service", handlers.TestHandle)
	database.Connect()

	// listening for requests
	fmt.Println("server is running")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
