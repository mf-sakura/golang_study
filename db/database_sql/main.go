package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mf-sakura/golang_study/db/database_sql/infrastructure"
	"github.com/mf-sakura/golang_study/db/database_sql/interfaces/controllers"
)

func main() {
	sqlHandler := infrastructure.NewSQLHandler()
	userController := controllers.NewUserController(sqlHandler.Conn)
	// フラグを-a追加その後も同様
	option := flag.String("a", "-h", "action")
	id := flag.String("i", "", "user id")
	firstName := flag.String("f", "Alan", "first name")
	lastName := flag.String("l", "Turing", "last name")
	flag.Parse()

	switch *option {
	// ユーザー一覧
	case "index":
		users, err := userController.Index()
		if err != nil {
			log.Fatal(err)
		}
		for _, user := range users {
			fmt.Printf("ID: %v, FirstName: %v, LastName: %v\n", user.ID, user.FirstName, user.LastName)
		}
		return
	// ユーザー詳細
	case "show":
		if *id == "" {
			log.Fatal("You need a user.id.")
		}
		user, err := userController.Show(*id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %v, FirstName: %v, LastName: %v\n", user.ID, user.FirstName, user.LastName)
		return
	// ユーザー作成
	case "create":
		if *firstName == "" {
			log.Fatal("You need a first name.")
		}
		if *lastName == "" {
			log.Fatal("You need a last name.")
		}
		user, err := userController.Create(*firstName, *lastName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %v, FirstName: %v, LastName: %v\n", user.ID, user.FirstName, user.LastName)
		return
	}
}
