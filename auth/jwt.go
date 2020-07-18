package auth

import (
	"fmt"
	"net/http"
	"strings"

	session "github.com/afifialaa/USER-AUTH/sessions"
	"github.com/dgrijalva/jwt-go"
)

// Generate token
func GenerateToken(email string) string {
	secretKey := []byte("secret key")

	fmt.Println("#generate token")
	// Create a token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	// Sign and get complete encoded token as string using the secret
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		fmt.Println("creating token error")
	}

	return tokenString
}

// Get token from request
func GetToken(r *http.Request) string {

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	if reqToken != "" {
		fmt.Println("#Gettoken -> token from request is: ", reqToken)
		return reqToken
	}

	return ""
}

// Validate token
// Sould be a middleware (i think)
func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return token, nil
	})

	if err != nil {
		fmt.Println("there was an error but we ignore it")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		var loggedUser string = session.GetLoggedUser()
		if claims["email"] == loggedUser {
			fmt.Println("claims are correct")
			return true
		} else {
			fmt.Println("claims are incorrect")
			return false
		}
	} else {
		return false
	}

}
