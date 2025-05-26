package auth

import "fmt"

// If we want to automatically export this funct to main,Function should be starts from capital letter
// if you want make as private then function should start from  small letter
func LoginWithCredentials(username string, password string) {
	fmt.Println("Login with Credentials", username, password)
}
