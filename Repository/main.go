package main

/*
	- Repository design pattern
	- Author : Okky Muhamad Budiman
	- Email : budimanokky93@gmail.com
*/

import (
	"fmt"
	"okky/go-design-pattern/Repository/handlers"
	"okky/go-design-pattern/Repository/singleton"
)

/*
	- Repository Pattern
	- Repository Pattern based on domain-driven-design
	- Also we implement Singleton pattern for handling database
	- Assume we have User domain in app
	- We create handler for user and repository for user
*/

func main() {
	// initial database
	dbHandler := singleton.NewDatabaseInitialize("root", "password", "repository")
	db, err := dbHandler.GetConnection()

	if err != nil {
		fmt.Println("error connect database : " + err.Error())
	}

	// initial object user handler
	userCtrl := handlers.NewUserController(db)
	userCtrl.GetUserByID()
}
