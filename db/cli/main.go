package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mf-sakura/golang_study/db/cli/infrastructure"
	"github.com/mf-sakura/golang_study/db/cli/interfaces/controllers"
)

func main() {
	sqlHandler := infrastructure.NewSQLHandler()
	userController := controllers.NewUserController(sqlHandler.Conn)
	option := flag.String("a", "", "action")
	id := flag.String("i", "", "user id")
	firstName := flag.String("f", "Alan", "first name")
	lastName := flag.String("l", "Turing", "last name")
	flag.Parse()

	switch *option {
	case "index":
		users, err := userController.Index()
		if err != nil {
			log.Fatal(err)
		}
		if len(users) == 0 {
			fmt.Println("No users found.")
		}
		for _, user := range users {
			fmt.Printf("ID: %v, FirstName: %v, LastName: %v\n", user.ID, user.FirstName, user.LastName)
		}
		return
	case "update":
		if *id == "" {
			log.Fatal("You need a user.id.")
		}
		if *firstName == "" {
			log.Fatal("You need a first name.")
		}
		if *lastName == "" {
			log.Fatal("You need a last name.")
		}
		err := userController.Update(*id, *firstName, *lastName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %v, FirstName: %v, LastName: %v\n", *id, *firstName, *lastName)
		return
	}
}
