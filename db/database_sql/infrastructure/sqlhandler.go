package infrastructure

import (
	"database/sql"
	"log"
	"os"

	// mysql driver
	// driverを直接参照する必要はなくinit関数を呼びたいだけなので、blank importをする。
	// https://golang.org/doc/effective_go.html#blank_import
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// SQLHandler is a struct for db connection
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler make SQLHandler
func NewSQLHandler() *SQLHandler {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	// `.env`を読み込む
	// RubyのdotenvのGo版
	err := godotenv.Load()
	if err != nil {
		// Fatalでプロセスが終了する
		log.Fatalf("error loading .env file. %s", err)
	}

	// 環境変数の読み込み
	database := os.Getenv("DATABASE_DATASOURCE")
	conn, err := sql.Open("mysql", database)
	if err != nil {
		log.Fatalf("error opening mysql. %s", err)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
