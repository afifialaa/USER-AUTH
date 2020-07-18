package session

import (
	  "os"
)

// Starts a new Session
func Start(email string){
	os.Setenv("userEmail", email)
}

// Returns logged in user
func GetLoggedUser() string{
	  if os.Getenv("userEmail") == ""{
		    return ""
	  }else{
		    return os.Getenv("userEmail")
	  }
}

// Ends user session
func End(){
	  os.Unsetenv("userEmail")
}
