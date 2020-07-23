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
	bearerToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearerToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	} else {
		return ""
	}

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

	fmt.Println(claims["email"])
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

func VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret key"), nil
	})
	fmt.Println(token.email)
	if err != nil {
		return false
	}
	return true
}


func Test(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["email"])
		return true
	} else {
		fmt.Println(err)
		return false
	}
}
