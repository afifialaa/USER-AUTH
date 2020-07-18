package validation

import (
	"fmt"
	"regexp"

	"github.com/afifialaa/USER-AUTH/helpers"
)

type User_type struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type User_login_type struct {
	Email    string
	Password string
}

func ValidateUser(user *User_type) bool {

	//empty fields
	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
		fmt.Println("empyty fields")
		return false
	}

	//empty field
	if len(user.FirstName) == 0 || len(user.LastName) == 0 || len(user.Email) == 0 || len(user.Password) == 0 {
		fmt.Println("empty fields")

		return false
	}

	//validating email
	if !validateEmail(user.Email) {
		fmt.Println("not a valid email")
		return false
	}

	fmt.Println("email: "+user.Email+" ", validateEmail(user.Email))

	user.Password = helpers.HashPassword(user.Password)
	return true
}

func ValidateUserLogin(user *User_login_type) bool {
	return true
}

// Validating email
func validateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}
