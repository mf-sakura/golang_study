package infrastructure

import (
	"log"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/joho/godotenv"
)

// SQLHandler is a struct for db connection
type SQLHandler struct {
	Conn *sqlx.DB
}

// NewSQLHandler make SQLHandler
func NewSQLHandler() *SQLHandler {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file. %s", err)
	}

	database := os.Getenv("DATABASE_DATASOURCE")
	conn, err := sqlx.Open("mysql", database)
	if err != nil {
		log.Fatalf("error opening mysql. %s", err)
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
