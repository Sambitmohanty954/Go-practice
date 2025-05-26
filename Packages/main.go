package main

import (
	"github.com/Sambitmohanty954/Go-practice/auth"
)

//	go mod tidy
// go get github.com/Sambitmohanty954/Go-practice/

func main() {
	auth.LoginWithCredentials("admin", "admin")
	auth.GetSession("Sambitmohanty954")
}
