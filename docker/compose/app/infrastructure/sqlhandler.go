package infrastructure

import (
	"log"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/kelseyhightower/envconfig"
)

// SQLHandler is a struct for db connection
type SQLHandler struct {
	Conn *sqlx.DB
}

// SQLConfig is config for sql connection
type SQLConfig struct {
	DataSourceName string `envconfig:"DATABASE_DATASOURCE"`
}

// NewSQLHandler make SQLHandler
func NewSQLHandler() *SQLHandler {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	var config SQLConfig
	if err := envconfig.Process("", &config); err != nil {
		log.Fatal(err)
	}

	conn, err := sqlx.Open("mysql", config.DataSourceName)
	if err != nil {
		panic(err.Error)
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
