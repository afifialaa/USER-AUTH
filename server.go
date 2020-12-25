package main

import (
	"fmt"
	"net/http"

	database "github.com/afifialaa/USER-AUTH/database"
	"github.com/afifialaa/USER-AUTH/handlers"
)

type Status struct {
	msg string
}

func main() {
	database.MongoConnect()

	// routes
	http.HandleFunc("/user/createUser", handlers.Signup)
	http.HandleFunc("/user/signin", handlers.Login)
	http.HandleFunc("/api/service/test", handlers.TestHandle)

	// listening for requests
	fmt.Println("server is running")
	http.ListenAndServe(":8080", nil)

}
