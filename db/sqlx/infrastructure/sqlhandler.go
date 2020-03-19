package infrastructure

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type SqlHandler struct {
	Conn *sqlx.DB
}

func NewSqlHandler() *SqlHandler {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file. %s", err)
	}

	database := os.Getenv("DATABASE_DATASOURCE")
	conn, err := sqlx.Open("mysql", database)
	if err != nil {
		panic(err.Error)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
